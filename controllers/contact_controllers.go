// controllers/contact_controller.go
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"

	"go-contact-api/database"
	"go-contact-api/models"
	"github.com/nyaruka/phonenumbers"
)

func CreateContact(c *gin.Context) {
	var input models.Contact
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var phones []models.Phone
	for _, number := range input.PhoneNumbers {
		parsed, _ := phonenumbers.Parse(number, "AU")
		e164 := phonenumbers.Format(parsed, phonenumbers.E164)
		phones = append(phones, models.Phone{Number: e164})
	}
	input.Phones = phones

	if err := database.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}
