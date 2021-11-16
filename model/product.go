package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name  string  `json:"name"`
	Price float64 `json:"price"`

	OwnerID uint `json:"owner_id"`
	Owner   User `json:"owner" gorm:"foreignKey:OwnerID;references:ID"`
}
