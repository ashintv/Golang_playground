package controllers

import (
	"crud_api/configs"
	"crud_api/models"
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