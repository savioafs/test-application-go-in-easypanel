package main

import (
	"net/http"

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

}
