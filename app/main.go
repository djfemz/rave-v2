package main

import (
	"github.com/djfemz/rave/app/security/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/login", controllers.NewAuthController().LoginHandler)

	err := router.Run(":8082")
	if err != nil {
		log.Println("Error starting server: ", err)
	}
}
