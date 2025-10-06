package models

import "gorm.io/gorm"

// user table

type Users struct {
	gorm.Model
	Username string
	Password string
	Post    []Post `gorm:"foreignKey:UserID"`
}

// posts table
type Post struct {
	gorm.Model
	Title   string
	Content string
	UserID    uint
}