package service

import (
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/repository"
	"gorm.io/gorm"
)

type ProductService struct {
	ProductRepo *repository.ProductRespository
}

func (c *ProductService) StoreProduct(product *model.Product) error {
	if err := c.ProductRepo.Store(product); err != nil {
		return err
	}

	return nil
}

func (c *ProductService) GetProduct(product_id *uint) (*model.Product, error) {
	product, err := c.ProductRepo.FindById(product_id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (c *ProductService) ListProducts() (*[]model.Product, error) {
	products, err := c.ProductRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (c *ProductService) ListUserProducts(owner_id *uint) (*[]model.Product, error) {
	products, err := c.ProductRepo.FindAllOwnerByUser(owner_id)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (c *ProductService) UpdateProduct(product *model.Product, product_id *uint) error {
	if !c.ProductRepo.IsExists(product_id) {
		return gorm.ErrRecordNotFound
	}
	if err := c.ProductRepo.Update(product_id, product); err != nil {
		return err
	}
	new_product, err := c.ProductRepo.FindById(product_id)
	if err != nil {
		return err
	}
	*product = *new_product

	return nil
}

func (c *ProductService) DeleteProduct(product_id *uint) error {
	if !c.ProductRepo.IsExists(product_id) {
		return gorm.ErrRecordNotFound
	}
	if err := c.ProductRepo.Delete(product_id); err != nil {
		return err
	}

	return nil
}
