package model

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID       int `gorm:"primaryKey"`
	Isbn     string
	Title    string
	AuthorID uint `gorm:"foreignKey:AuthorID;references:ID"`
	Author   Author
}

type Author struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
}
