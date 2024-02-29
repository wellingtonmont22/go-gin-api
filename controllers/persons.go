package controllers

import (
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindPersons(c *gin.Context) {
	var persons []models.Person

	models.DB.Find(&persons)

	c.JSON(http.StatusOK, gin.H{"data": persons})
}
