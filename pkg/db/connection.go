package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() {
	gorm.Open(postgres.Open(""))
}
