package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

type City struct {
	ID          int    `json:"id,omitempty"  db:"ID"`
	Name        string `json:"name,omitempty"  db:"Name"`
	CountryCode string `json:"countryCode,omitempty"  db:"CountryCode"`
	District    string `json:"district,omitempty"  db:"District"`
	Population  int    `json:"population,omitempty"  db:"Population"`
}

var (
	db *sqlx.DB
)

func main() {
	_db, err := sqlx.Open("mysql", "root:password@tcp(db:3306)/world?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}
	db = _db

	fmt.Println("Connected!")
	var city City
	if err := db.Get(&city, "SELECT * FROM city WHERE Name='Tokyo'"); errors.Is(err, sql.ErrNoRows) {
		log.Printf("no such city Name = %s", "Tokyo")
	} else if err != nil {
		log.Fatalf("DB Error: %s", err)
	}
	fmt.Printf("Tokyoの人口は%d人です\n", city.Population)

	// cities := []City{}
	// db.Select(&cities, "SELECT * FROM city WHERE CountryCode='JPN'")
	// fmt.Println("日本の都市一覧")
	// for _, city := range cities {
	// 	fmt.Printf("都市名: %s, 人口: %d人\n", city.Name, city.Population)
	// }
	e := echo.New()

	e.GET("/cities/:cityName", getCityInfoHandler)
	e.POST("/cities", postCityInfoHandler)

	e.Start(":8080")
}

func getCityInfoHandler(c echo.Context) error {
	cityName := c.Param("cityName")
	fmt.Println(cityName)

	var city City
	if err := db.Get(&city, "SELECT * FROM city WHERE Name=?", cityName); errors.Is(err, sql.ErrNoRows) {
		log.Printf("No Such City Name=%s", cityName)
	} else if err != nil {
		log.Fatalf("DB Error: %s", err)
	}

	return c.JSON(http.StatusOK, city)
}

func postCityInfoHandler(c echo.Context) error {
	var city City
	if err := c.Bind(&city); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	ctx := context.Background()
	_, err := db.ExecContext(ctx, "INSERT INTO `city` VALUES (?,?,?,?,?)", city.ID, city.Name, city.CountryCode, city.District, city.Population) 
  if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, city)
}
