package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int    `gorm:"PrimaryKey"`
	Username string `gorm:"unique"`
}
