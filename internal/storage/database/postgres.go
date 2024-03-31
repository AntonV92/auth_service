package database

import (
	"auth_service/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func New(dbConf config.DbConfig) (*sql.DB, error) {

	source := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConf.Host, dbConf.Port, dbConf.User, dbConf.Pass, dbConf.DbName, dbConf.SSLMode)

	db, err := sql.Open("postgres", source)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
