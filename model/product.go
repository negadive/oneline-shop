package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name    string `json:"name"`
	OwnerID uint   `json:"owner_id"`
	Owner   User   `json:"owner" gorm:"foreignKey:OwnerID;references:ID"`
}
