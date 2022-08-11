package db

import (
	"github.com/jinzhu/gorm"
	"gorm.io/driver/postgres"
)

dsn = "host=localhost user=postgres"

func Connection() {
	gorm.Open(postgres.Open())
}
