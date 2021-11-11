package schema

import "time"

/*************
REQUEST SCHEMA
*************/
type UserRegisterReq struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,min=4,max=100"`
	Password string `json:"password" validate:"required,min=6,max=45"`
}

type UserUpdateReq struct {
	Name     string `json:"name" validate:"min=4,max=100"`
	Password string `json:"password" validate:"min=6,max=45"`
}

/**************
RESPONSE SCHEMA
**************/
type UserRegisterRes struct {
	Email     string    `json:"email"`
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type UserUpdateRes struct {
	Email     string    `json:"email"`
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
}
