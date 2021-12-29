package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email    string `json:"email" gorm:"unique,not null"`
	Name     string `json:"name" gorm:"unique,not null"`
	Password string `json:"password" gorm:"unique,not null"`
	// Products []Product
}
