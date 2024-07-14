package main

import (
	"GIN-GORM-Tutorial/controllers"
	"GIN-GORM-Tutorial/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/posts", controllers.PostsCreate)
	r.POST("/posts", controllers.PostsCreate)
	r.Run()
}
