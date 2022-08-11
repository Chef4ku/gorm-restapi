package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title       string `gorm:"type:varchar(255);not null;unique_index"`
	Description string
	UserID      uint
	Done        bool `gorm:"default:false"`
}
