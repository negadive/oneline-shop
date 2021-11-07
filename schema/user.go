package schema

type UserRegisterReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserRegisterRes struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
