package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string `gorm:"unique"`
	gmail    string `gorm:"unique"`
	age      int    `gorm:"not null"`
}
