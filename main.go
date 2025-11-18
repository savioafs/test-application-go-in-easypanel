package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/savioafs/test-application-go-in-easypanel/internal/controller"
	"github.com/savioafs/test-application-go-in-easypanel/internal/database"
	"github.com/savioafs/test-application-go-in-easypanel/internal/entity"
)

func main() {

	database.Connect()

	database.DB.AutoMigrate(&entity.User{})

	r := gin.Default()

	r.POST("/client", controller.CreateClient)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on :" + port)
	if err := r.Run("0.0.0.0:" + port); err != nil {
		log.Fatal(err)
	}

}
