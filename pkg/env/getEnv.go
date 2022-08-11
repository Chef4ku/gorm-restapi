package env

import (
	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	err := godotenv.Load("./pkg/env/.env")
	if err != nil {
		panic("Could not read .env file. Make sure the .env file is created in pkg/env")
	}
}
