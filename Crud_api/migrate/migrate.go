package main

import (
	"crud_api/configs"
	"crud_api/models"
)

func main() {
	configs.LoadConfig()
	configs.DB.AutoMigrate(&models.Users{}, &models.Post{})
}
