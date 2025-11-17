package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/route-one", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "rota 01",
		})
	})

	r.GET("/route-two", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "rota 02",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on :" + port)
	if err := r.Run("0.0.0.0:" + port); err != nil {
		log.Fatal(err)
	}
}
