package controllers

import (
	"GIN-GORM-Tutorial/initializers"
	"GIN-GORM-Tutorial/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data of req body

	// Create a post
	post := models.Post{Title: "post title", Body: "post body"}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}
