package main

import "log"

func DeleteUser(id int) {
	query := "DELETE FROM users WHERE id=$1"

	_, err := db.Exec(query, id)
	if err != nil {
		log.Fatal("Fail : ", err)
	}

	log.Printf("Kullanıcı (ID: %d) başarıyla silindi.\n", id)
}
