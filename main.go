package main

import (
	"github.com/gin-gonic/gin"
	"go-contact-api/controllers"
	"go-contact-api/database"
	"go-contact-api/validators"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()
	validators.RegisterPhoneValidator()

	r.POST("/contacts", controllers.CreateContact)

	r.Run()
}
