package controller

import (
	"github.com/negadive/oneline/db"
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/schema"
	"gorm.io/gorm"
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

func GetProduct(product_id int) (*model.Product, error) {
	_db := db.GetDb()

	product := model.Product{}
	result := _db.First(&product, product_id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func listProducts(owner_id int) (*[]model.Product, error) {
	_db := db.GetDb()

	products := []model.Product{}
	query := _db.Model(&model.Product{})
	if owner_id != 0 {
		query = query.Where("owner_id = ?", owner_id)
	}

	result := query.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return &products, nil
}

func ListProducts() (*[]model.Product, error) {
	return listProducts(0)
}

func ListUserProducts(owner_id int) (*[]model.Product, error) {
	return listProducts(owner_id)
}

func UpdateProduct(_product *schema.ProductUpdateReq, product_id int) (*model.Product, error) {
	var count int64
	_db := db.GetDb()

	if _db.Model(&model.Product{}).Where("id = ?", product_id).Count(&count); count < 1 {
		return nil, gorm.ErrRecordNotFound
	}

	product := model.Product{}
	if err := _db.Model(&model.Product{}).Where("id = ?", product_id).Updates(model.Product{Name: _product.Name}).Error; err != nil {
		return nil, err
	}
	if err := _db.First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func DeleteProduct(product_id int) error {
	var count int64
	_db := db.GetDb()

	if _db.Model(&model.Product{}).Where("id = ?", product_id).Count(&count); count < 1 {
		return gorm.ErrRecordNotFound
	}

	if err := _db.Delete(&model.Product{}, product_id).Error; err != nil {
		return err
	}

	return nil
}
