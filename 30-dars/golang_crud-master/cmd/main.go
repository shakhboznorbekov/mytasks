package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/asadbekGo/golang_crud/api"
	"github.com/asadbekGo/golang_crud/config"
	"github.com/asadbekGo/golang_crud/pkg/db"
)

func main() {

	cfg := config.Load()

	db, err := db.ConnectionDB(&cfg)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	r := gin.New()

	api.SetUpApi(r, db)

	log.Printf("Listening port %v...\n", cfg.HTTPPort)
	err = r.Run(cfg.HTTPPort)
	if err != nil {
		panic(err)
	}
}
