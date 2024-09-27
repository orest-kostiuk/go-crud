package main

import (
	"github.com/gin-gonic/gin"
	"github.com/orest-kostiuk/go-crud/controllers"
	"github.com/orest-kostiuk/go-crud/initializers"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)

	if err := r.Run(); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
