package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email    string `json:"email" gorm:"unique,not null"`
	Name     string `json:"name"`
	Password string `json:"password"`
	// Products []Product
}
