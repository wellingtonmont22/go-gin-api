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

	r.Run() // listen and serve on 0.0.0.0:8080
}
