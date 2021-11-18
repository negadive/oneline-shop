package repository

import (
	"github.com/negadive/oneline/model"
	"gorm.io/gorm"
)

type ProductRespository struct {
	DBCon *gorm.DB
}

func (repo *ProductRespository) Store(product *model.Product) error {
	result := repo.DBCon.Create(&product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *ProductRespository) Update(product_id *uint, product *model.Product) error {
	if err := repo.DBCon.Model(&model.Product{}).Where("id = ?", *product_id).Updates(&product).Error; err != nil {
		return err
	}

	return nil
}

func (repo *ProductRespository) Delete(product_id *uint) error {
	if err := repo.DBCon.Delete(&model.Product{}, product_id).Error; err != nil {
		return err
	}

	return nil
}

func (repo *ProductRespository) FindById(product_id *uint) (*model.Product, error) {
	product := model.Product{}
	result := repo.DBCon.First(&product, product_id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (repo *ProductRespository) FindAll() (*[]model.Product, error) {
	products := []model.Product{}
	result := repo.DBCon.Model(&model.Product{}).Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return &products, nil
}

func (repo *ProductRespository) FindAllOwnerByUser(owner_id *uint) (*[]model.Product, error) {
	products := []model.Product{}
	result := repo.DBCon.Model(&model.Product{}).Where("owner_id = ?", owner_id).Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return &products, nil
}

func (repo *ProductRespository) IsExists(product_id *uint) bool {
	var count int64
	if repo.DBCon.Model(&model.Product{}).Where("id = ?", product_id).Count(&count); count < 1 {
		return false
	}

	return count > 0

}

func (repo *ProductRespository) ProductWithOwnerExists(product_id *uint, owner_id *uint) bool {
	var count int64

	repo.DBCon.Model(&model.Product{}).Where("id = ?", *product_id).Where("owner_id = ?", *owner_id).Count(&count)

	return count == 1
}

func (repo *ProductRespository) FindByIds(products *[]model.Product, product_ids *[]uint) error {
	if err := repo.DBCon.Model(&model.Product{}).Find(products, product_ids).Error; err != nil {
		return err
	}

	return nil
}
