package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/savioafs/test-application-go-in-easypanel/internal/database"
	"github.com/savioafs/test-application-go-in-easypanel/internal/entity"
)

func CreateClient(c *gin.Context) {
	var input struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := entity.NewUser(input.Name, input.Email)

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "não foi possível salvar o usuário",
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}
