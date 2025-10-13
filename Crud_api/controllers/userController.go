package controllers

import (
	"crud_api/configs"
	"crud_api/models"
	"log"
	"github.com/gin-gonic/gin"
)


type createUserRequest struct{
	Username string `json:"username" binding:"required"`
    Password string `json:"Password" binding:"required"`
}
func AddUser(ctx *gin.Context) {
	var req createUserRequest
	log.Print("Add user")
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid data",
			"error":   err,
		})
		return
	}

	user := models.Users{
		Username: req.Username,
		Password: req.Password,
	}


	id := configs.DB.Create(&user)
	ctx.JSON(200, gin.H{
		"userId": id,
	})

}
