package schema

/*************
REQUEST SCHEMA
*************/
type LoginReq struct {
	Name     string `json:"name" validate:"required,min=4,max=100"`
	Password string `json:"password" validate:"required,min=6,max=45"`
}

/**************
RESPONSE SCHEMA
**************/
