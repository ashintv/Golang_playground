package controllers

import (
	"crud_api/configs"
	"crud_api/models"
	"log"

	"github.com/gin-gonic/gin"
)

func AddUser(ctx *gin.Context) {
	var user models.Users
	log.Print("Add user")
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid data",
			"error":   err,
		})
		return
	}

	// hash passs
	id := configs.DB.Create(&user)
	ctx.JSON(200, gin.H{
		"userId": id,
	})

}
