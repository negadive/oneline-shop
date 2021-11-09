package schema

import "time"

/*************
REQUEST SCHEMA
*************/
type ProductStoreReq struct {
	Name    string `json:"name" validate:"required,min=4,max=100"`
	OwnerID uint   `json:"owner_id" validate:"required"`
}

type ProductUpdateReq struct {
	Name string `json:"name" validate:"required,min=4,max=100"`
}

/**************
RESPONSE SCHEMA
**************/

type ProductStoreRes struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	OwnerID   uint      `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ProductGetOneRes struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	OwnerID   uint      `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductListRes struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	OwnerID uint   `json:"owner_id"`
}

type ProductUpdateRes struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	OwnerID   uint      `json:"owner_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
