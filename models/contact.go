package models

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	FullName     string   `json:"full_name" binding:"required"`
	Email        string   `json:"email" binding:"omitempty,email"`
	PhoneNumbers []string `json:"phone_numbers" gorm:"-" binding:"required,uniquePhoneNumbers"`
	Phones       []Phone  `json:"-" gorm:"foreignKey:ContactID"`
}

type Phone struct {
	gorm.Model
	ContactID uint
	Number    string `json:"number" binding:"required"`
}
