package main

import (
	"auth_service/config"
	"log"
)

func main() {

	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
}
