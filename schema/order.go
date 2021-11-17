package schema

import "time"

/*************
REQUEST SCHEMA
*************/
type OrderStoreReq struct {
	Note       string `json:"note" validate:"omitempty,max=150"`
	ProductIDs []uint `json:"product_ids" valdlidate:"required,min=1,max=20"`
}

/**************
RESPONSE SCHEMA
**************/
type OrderStoreRes struct {
	ID         uint      `json:"id"`
	TotalPrice float64   `json:"total_price"`
	CustomerID uint      `json:"customer_id"`
	Status     string    `json:"status"`
	PaidAt     time.Time `json:"paid_at"`
	Note       string    `json:"note"`
	CreatedAt  time.Time `json:"created_at"`
}

type OrderGetOneRes struct {
	ID         uint      `json:"id"`
	TotalPrice float64   `json:"total_price"`
	CustomerID uint      `json:"customer_id"`
	Status     string    `json:"status"`
	PaidAt     time.Time `json:"paid_at"`
	Note       string    `json:"note"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedAt  time.Time `json:"created_at"`
}

type OrderListRes struct {
	ID         uint    `json:"id"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"`
	Note       string  `json:"note"`
}

type OrderUpdateRes struct {
	ID         uint      `json:"id"`
	TotalPrice float64   `json:"total_price"`
	CustomerID uint      `json:"customer_id"`
	Status     string    `json:"status"`
	PaidAt     time.Time `json:"paid_at"`
	Note       string    `json:"note"`
	UpdatedAt  time.Time `json:"updated_at"`
}
