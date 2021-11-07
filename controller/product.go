package controller

import (
	"github.com/negadive/oneline/db"
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/schema"
)

func StoreProduct(_product *schema.ProductStoreReq) (*model.Product, error) {
	_db := db.GetDb()

	product := model.Product{Name: _product.Name, OwnerID: _product.OwnerID}
	result := _db.Create(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func ListProducts() (*[]model.Product, error) {
	_db := db.GetDb()

	products := []model.Product{}
	result := _db.Model(&model.Product{}).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return &products, nil
}
