package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type City struct {
	ID          int    `json:"id,omitempty"  db:"ID"`
	Name        string `json:"name,omitempty"  db:"Name"`
	CountryCode string `json:"countryCode,omitempty"  db:"CountryCode"`
	District    string `json:"district,omitempty"  db:"District"`
	Population  int    `json:"population,omitempty"  db:"Population"`
}

func main() {
	db, err := sqlx.Open("mysql", "root:password@tcp(local_mysql:3306)/world?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}

	fmt.Println("Connected!")
	// var city City
	// if err := db.Get(&city, "SELECT * FROM city WHERE Name='Tokyo'"); errors.Is(err, sql.ErrNoRows) {
	// 	log.Printf("no such city Name = %s", "Tokyo")
	// } else if err != nil {
	// 	log.Fatalf("DB Error: %s", err)
	// }

	// fmt.Printf("Tokyoの人口は%d人です\n", city.Population)
	err = db.Ping()
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}
}
