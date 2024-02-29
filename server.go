package main

import (
	"api/controllers"
	"api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/persons", controllers.FindPersons)
	r.POST("/persons", controllers.CreatePerson)
	r.PUT("/persons/:id", controllers.UpdatePerson)
	r.Run() // listen and serve on 0.0.0.0:8080
}
