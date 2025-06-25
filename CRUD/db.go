package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnDb() {
	var err error
	ConnStr := "host=localhost port=5432 user=postgres password=12345 dbname=kullanicilar sslmode=disable"
	db, err = sql.Open("postgres", ConnStr)
	if err != nil {
		log.Fatal("Connection Error: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Connection Fail: ", err)
	}

	fmt.Print("Succesfully Connection")
}
