package controllers

import (
	"crud_api/configs"
	"crud_api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)


type CreatePostRequest struct {
    Title   string `json:"title" binding:"required"`
    Content string `json:"content" binding:"required"`
}


func CreatePost(ctx *gin.Context){
	var req CreatePostRequest
	if err:= ctx.ShouldBindJSON(&req); err!=nil{
		ctx.JSON(404 , gin.H{
			"message":"invalid data",
			"err":err,
		})
		return
	}

	post := models.Post{
		Title: req.Title,
		Content: req.Content,
 	}


	res:= configs.DB.Create(&post)
	if res.Error!=nil{
		ctx.JSON(500 , gin.H{
			"message":"Db request failed",
		})
		return
	}
	ctx.JSON(200 , gin.H{
			"id":post.ID,
	})
}


func GetPost(ctx *gin.Context) {
	id := ctx.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Invalid user ID",
			"error":   err.Error(),
		})
		return
	}

	var post models.Post

	result := configs.DB.Where(&models.Post{UserID:uint(userID)}).First(&post)

	if result.Error != nil {
		ctx.JSON(404, gin.H{
			"message": "Post not found",
			"error":   result.Error.Error(),
		})
		return
	}

	ctx.JSON(200, post)
}
