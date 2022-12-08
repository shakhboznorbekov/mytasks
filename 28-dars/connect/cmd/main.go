package main

import (
	"app/config"
	"app/models"
	"app/pkg/db"
	"app/storage"
	"fmt"
	"log"
)

func main() {
	cfg := config.Load()

	db, err := db.ConnectionDB(&cfg)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	u := models.User{
		FirstName: "Shaxboz",
		LastName:  "Norbekov",
	}

	id, err := storage.Create(db, u)
	if err != nil {
		log.Fatalf("error whiling create user: %v", err)
	}
	fmt.Println(id)

	user, err := storage.GetById(db, id)
	if err != nil {
		log.Fatalf("error whiling create user : %v", err)
	}

	fmt.Println(user)

	// users, err := storage.GetList(db)
	// if err != nil {
	// 	log.Fatalf("error whiling create user : %v", err)
	// }

	// fmt.Println(users)
}
