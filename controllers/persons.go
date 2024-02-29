package controllers

import (
	"api/models"
	"api/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindPersons(c *gin.Context) {
	var persons []models.Person

	models.DB.Find(&persons)

	c.JSON(http.StatusOK, gin.H{"data": persons})
}

func CreatePerson(c *gin.Context) {
	var input schemas.CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	person := models.Person{FirstName: input.FirstName, LastName: input.LastName}

	models.DB.Create(&person)

	c.JSON(http.StatusOK, gin.H{"data": person})
}

func UpdatePerson(c *gin.Context) {
	var person models.Person

	if err := models.DB.Where("id = ?", c.Param("id")).First(&person).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input schemas.UpdateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&person).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": person})
}
