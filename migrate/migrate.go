package main

import (
	"github.com/orest-kostiuk/go-crud/initializers"
	"github.com/orest-kostiuk/go-crud/models"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// Perform AutoMigrate and handle any error
	err := initializers.DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
}
