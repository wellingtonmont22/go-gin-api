package controllers

import (
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

func FindPersons(c *gin.Context) {
	var persons []models.Person

	models.DB.Find(&persons)

	c.JSON(http.StatusOK, gin.H{"data": persons})
}

func CreatePerson(c *gin.Context) {
	var input CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person := models.Person{FirstName: input.FirstName, LastName: input.LastName}

	models.DB.Create(&person)

	c.JSON(http.StatusOK, gin.H{"data": person})
}
