package main

import (
	"auth_service/config"
	"auth_service/internal/storage/database"
	"log"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(config.DbConf)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
