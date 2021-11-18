package repository

import (
	"context"

	"github.com/negadive/oneline/model"
	"gorm.io/gorm"
)

type IProductRepository interface {
	Store(ctx context.Context, product *model.Product) error
	Update(ctx context.Context, product_id *uint, product *model.Product) error
	Delete(ctx context.Context, product_id *uint) error
	FindById(ctx context.Context, product_id *uint) (*model.Product, error)
	FindAll(ctx context.Context) (*[]model.Product, error)
	FindAllOwnerByUser(ctx context.Context, owner_id *uint) (*[]model.Product, error)
	IsExists(ctx context.Context, product_id *uint) bool
	ProductWithOwnerExists(ctx context.Context, product_id *uint, owner_id *uint) bool
	FindByIds(ctx context.Context, products *[]model.Product, product_ids *[]uint) error
}

type ProductRepository struct {
	DBCon *gorm.DB
}

func NewProductRepository(DBCon *gorm.DB) IProductRepository {
	r := ProductRepository{
		DBCon: DBCon,
	}

	return &r
}

func (repo *ProductRepository) Store(ctx context.Context, product *model.Product) error {
	result := repo.DBCon.WithContext(ctx).Create(&product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *ProductRepository) Update(ctx context.Context, product_id *uint, product *model.Product) error {
	if err := repo.DBCon.WithContext(ctx).Model(&model.Product{}).Where("id = ?", *product_id).Updates(&product).Error; err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepository) Delete(ctx context.Context, product_id *uint) error {
	if err := repo.DBCon.WithContext(ctx).Delete(&model.Product{}, product_id).Error; err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepository) FindById(ctx context.Context, product_id *uint) (*model.Product, error) {
	product := model.Product{}
	result := repo.DBCon.WithContext(ctx).First(&product, product_id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (repo *ProductRepository) FindAll(ctx context.Context) (*[]model.Product, error) {
	products := []model.Product{}
	result := repo.DBCon.WithContext(ctx).Model(&model.Product{}).Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return &products, nil
}

func (repo *ProductRepository) FindAllOwnerByUser(ctx context.Context, owner_id *uint) (*[]model.Product, error) {
	products := []model.Product{}
	result := repo.DBCon.WithContext(ctx).Model(&model.Product{}).Where("owner_id = ?", owner_id).Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return &products, nil
}

func (repo *ProductRepository) IsExists(ctx context.Context, product_id *uint) bool {
	var count int64
	if repo.DBCon.WithContext(ctx).Model(&model.Product{}).Where("id = ?", product_id).Count(&count); count < 1 {
		return false
	}

	return count > 0

}

func (repo *ProductRepository) ProductWithOwnerExists(ctx context.Context, product_id *uint, owner_id *uint) bool {
	var count int64

	repo.DBCon.WithContext(ctx).Model(&model.Product{}).Where("id = ?", *product_id).Where("owner_id = ?", *owner_id).Count(&count)

	return count == 1
}

func (repo *ProductRepository) FindByIds(ctx context.Context, products *[]model.Product, product_ids *[]uint) error {
	if err := repo.DBCon.WithContext(ctx).Model(&model.Product{}).Find(products, product_ids).Error; err != nil {
		return err
	}

	return nil
}
