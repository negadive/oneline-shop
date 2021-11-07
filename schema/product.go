package schema

type ProductStoreReq struct {
	Name    string `json:"name"`
	OwnerID uint   `json:"owner_id"`
}

type ProductStoreRes struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	OwnerID uint   `json:"owner_id"`
}

type ProductListRes struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	OwnerID uint   `json:"owner_id"`
}
