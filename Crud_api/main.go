package main

import (
	"crud_api/configs"
	"crud_api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	configs.LoadConfig()
	router.POST("/user", controllers.AddUser)
	router.Run()
}
