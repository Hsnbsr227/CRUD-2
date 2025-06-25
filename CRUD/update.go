package main

import (
	"fmt"
	"log"
)

func UpdateUser(id int, newName string, newPassword string) {
	query := "UPDATE users SET username=$1, password =$2 WHERE id=$3"

	_, err := db.Exec(query, newName, newPassword, id)
	if err != nil {
		log.Fatal("User could not be updated:",err)
	}
	
	fmt.Printf("Kullanıcı (ID: %d) başarıyla güncellendi.\n", id)
}
