package models

import "gorm.io/gorm"

// user table

type Users struct {
	gorm.Model
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Post     []Post `gorm:"foreignKey:UserID"`
}

// posts table
type Post struct {
	gorm.Model
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserID  uint `json:"userId" binding:"required"`
}
