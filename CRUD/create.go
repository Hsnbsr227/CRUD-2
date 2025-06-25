package main

import (
	"fmt"
	"log"
)

func CreateUser(name string, password string) {
	query := "INSERT INTO users (username,password) VALUES ($1,$2)"
	_, err := db.Exec(query, name, password)
	if err != nil {
		log.Fatal("User could not be added: ", err)
	}
	fmt.Print("User Added: ", name)
}
