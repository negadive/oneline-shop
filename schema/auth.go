package schema

/*************
REQUEST SCHEMA
*************/
type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=45"`
}

/**************
RESPONSE SCHEMA
**************/
