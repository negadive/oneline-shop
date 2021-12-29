package repository

import (
	"context"

	"github.com/negadive/oneline/customErrors"
	"github.com/negadive/oneline/model"
	"gorm.io/gorm"
)

type IProductRepository interface {
	Store(ctx context.Context, tx *gorm.DB, product *model.Product) error
	Update(ctx context.Context, tx *gorm.DB, productId *uint, product *model.Product) error
	Delete(ctx context.Context, tx *gorm.DB, productId *uint) error
	FindById(ctx context.Context, tx *gorm.DB, productId *uint) (*model.Product, error)
	FindAll(ctx context.Context, tx *gorm.DB) (*[]*model.Product, error)
	FindAllOwnerByUser(ctx context.Context, tx *gorm.DB, ownerId *uint) (*[]*model.Product, error)
	IsExists(ctx context.Context, tx *gorm.DB, productId *uint) bool
	ProductWithOwnerExists(ctx context.Context, tx *gorm.DB, productId *uint, ownerId *uint) bool
	FindByIds(ctx context.Context, tx *gorm.DB, products *[]model.Product, productIds *[]uint) error
}

type ProductRepository struct {
}

func NewProductRepository() IProductRepository {
	return &ProductRepository{}
}

func (repo *ProductRepository) Store(ctx context.Context, tx *gorm.DB, product *model.Product) error {
	result := tx.WithContext(ctx).Create(&product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *ProductRepository) Update(ctx context.Context, tx *gorm.DB, productId *uint, product *model.Product) error {
	if err := tx.WithContext(ctx).Model(&model.Product{}).Where("id = ?", *productId).Updates(&product).Error; err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepository) Delete(ctx context.Context, tx *gorm.DB, productId *uint) error {
	if err := tx.WithContext(ctx).Delete(&model.Product{}, productId).Error; err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepository) FindById(ctx context.Context, tx *gorm.DB, productId *uint) (*model.Product, error) {
	product := model.Product{}
	result := tx.WithContext(ctx).First(&product, productId)
	if result.Error != nil {
		return nil, customErrors.NewNotFoundError("product")
	}

	return &product, nil
}

func (repo *ProductRepository) FindAll(ctx context.Context, tx *gorm.DB) (*[]*model.Product, error) {
	products := []*model.Product{}
	result := tx.WithContext(ctx).Model(&model.Product{}).Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return &products, nil
}

func (repo *ProductRepository) FindAllOwnerByUser(ctx context.Context, tx *gorm.DB, ownerId *uint) (*[]*model.Product, error) {
	products := []*model.Product{}
	result := tx.WithContext(ctx).Model(&model.Product{}).Where("owner_id = ?", ownerId).Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return &products, nil
}

func (repo *ProductRepository) IsExists(ctx context.Context, tx *gorm.DB, productId *uint) bool {
	var count int64
	if tx.WithContext(ctx).Model(&model.Product{}).Where("id = ?", productId).Count(&count); count < 1 {
		return false
	}

	return count > 0

}

func (repo *ProductRepository) ProductWithOwnerExists(ctx context.Context, tx *gorm.DB, productId *uint, ownerId *uint) bool {
	var count int64

	tx.WithContext(ctx).Model(&model.Product{}).Where("id = ?", *productId).Where("owner_id = ?", *ownerId).Count(&count)

	return count == 1
}

func (repo *ProductRepository) FindByIds(ctx context.Context, tx *gorm.DB, products *[]model.Product, productIds *[]uint) error {
	if err := tx.WithContext(ctx).Model(&model.Product{}).Find(products, productIds).Error; err != nil {
		return err
	}

	return nil
}
