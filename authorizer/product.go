package authorizer

import (
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/repository"
)

type IProductAuthorizer interface {
	CanCreate(product *model.Product, user_id *uint) bool
	CanEdit(product *model.Product, user_id *uint) bool
	CanDelete(product *model.Product, user_id *uint) bool
}

type ProductAuthorizer struct {
	productRepo repository.IProductRepository
}

func NewProductAuthorizer(productRepo repository.IProductRepository) IProductAuthorizer {
	return &ProductAuthorizer{productRepo: productRepo}
}

func (pa *ProductAuthorizer) CanCreate(product *model.Product, user_id *uint) bool {
	return product.OwnerID == *user_id
}
func (pa *ProductAuthorizer) CanEdit(product *model.Product, user_id *uint) bool {
	return product.OwnerID == *user_id
}

func (pa *ProductAuthorizer) CanDelete(product *model.Product, user_id *uint) bool {
	return product.OwnerID == *user_id
}
