package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	TotalPrice float64        `json:"total_price"`
	CustomerID uint           `json:"customer_id"`
	Customer   User           `json:"customer" gorm:"foreignKey:CustomerID;references:ID"`
	Status     string         `json:"status"`
	PaidAt     time.Time      `json:"paid_at"`
	Note       sql.NullString `json:"note"`
}

type OrderProduct struct {
	gorm.Model
	OrderID     uint    `json:"order_id,omitempty"`
	Order       Order   `json:"order,omitempty"`
	ProductID   uint    `json:"product_id,omitempty"`
	ProductName string  `json:"product_name,omitempty"`
	Product     Product `json:"product,omitempty"`
	Price       float64 `json:"price,omitempty"`
}
