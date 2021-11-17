package repository

import (
	"github.com/negadive/oneline/model"
	"gorm.io/gorm"
)

type ProductRespository struct {
	DBCon *gorm.DB
}

func (repo *ProductRespository) ProductWithOwnerExists(product_id int, owner_id uint) bool {
	var count int64

	repo.DBCon.Model(&model.Product{}).Where("id = ?", product_id).Where("owner_id = ?", owner_id).Count(&count)

	return bool(count == 1)
}

func (repo *ProductRespository) FindByIds(products *[]model.Product, product_ids *[]uint) error {
	if err := repo.DBCon.Model(&model.Product{}).Find(products, product_ids).Error; err != nil {
		return err
	}

	return nil
}
