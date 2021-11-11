package service

import (
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/schema"
	"gorm.io/gorm"
)

type ProductService struct {
	DBCon *gorm.DB
}

func (c *ProductService) StoreProduct(_product *schema.ProductStoreReq) (*model.Product, error) {
	product := model.Product{Name: _product.Name, OwnerID: _product.OwnerID}
	result := c.DBCon.Create(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (c *ProductService) GetProduct(product_id int) (*model.Product, error) {
	product := model.Product{}
	result := c.DBCon.First(&product, product_id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (c *ProductService) listProducts(owner_id int) (*[]model.Product, error) {
	products := []model.Product{}
	query := c.DBCon.Model(&model.Product{})
	if owner_id != 0 {
		query = query.Where("owner_id = ?", owner_id)
	}

	result := query.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return &products, nil
}

func (c *ProductService) ListProducts() (*[]model.Product, error) {
	return c.listProducts(0)
}

func (c *ProductService) ListUserProducts(owner_id int) (*[]model.Product, error) {
	return c.listProducts(owner_id)
}

func (c *ProductService) UpdateProduct(_product *schema.ProductUpdateReq, product_id int) (*model.Product, error) {
	var count int64
	if c.DBCon.Model(&model.Product{}).Where("id = ?", product_id).Count(&count); count < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	product := model.Product{}
	if err := c.DBCon.Model(&model.Product{}).Where("id = ?", product_id).Updates(model.Product{Name: _product.Name}).Error; err != nil {
		return nil, err
	}
	if err := c.DBCon.Where("id = ?", product_id).First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (c *ProductService) DeleteProduct(product_id int) error {
	var count int64
	if c.DBCon.Model(&model.Product{}).Where("id = ?", product_id).Count(&count); count < 1 {
		return gorm.ErrRecordNotFound
	}

	if err := c.DBCon.Delete(&model.Product{}, product_id).Error; err != nil {
		return err
	}

	return nil
}
