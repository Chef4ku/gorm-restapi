package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := getDSN()

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func getDSN() string {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	user := os.Getenv("USER")
	if user == "" {
		panic("Please set an environment variable with the key \"USER\" that is not empty in pkg/env/.env")
	}
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	if dbname == "" {
		panic("Please set an environment variable with the key \"DBNAME\" that is not empty, in pkg/env/.env")
	}
	port := os.Getenv("PORT")
	if port == "" {
		panic("Please set an environment variable with the key \"PORT\" that is not empty in pkg/env/.env")
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)
}
