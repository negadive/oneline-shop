package repository

import (
	"context"

	"github.com/negadive/oneline/model"
	"gorm.io/gorm"
)

type IOrderRepository interface {
	Store(ctx context.Context, tx *gorm.DB, order *model.Order) error
	StoreOrderProducts(ctx context.Context, tx *gorm.DB, order_id *uint, products *[]model.Product) error
}

type OrderRepository struct{}

func NewOrderRepository() IOrderRepository {
	r := OrderRepository{}

	return &r
}

func (r *OrderRepository) Store(ctx context.Context, tx *gorm.DB, order *model.Order) error {
	result := tx.WithContext(ctx).Create(order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *OrderRepository) StoreOrderProducts(ctx context.Context, tx *gorm.DB, order_id *uint, products *[]model.Product) error {
	order_products := []model.OrderProduct{}
	for _, product := range *products {
		order_product := model.OrderProduct{
			OrderID:     *order_id,
			ProductID:   product.ID,
			ProductName: product.Name,
			Price:       product.Price,
		}

		order_products = append(order_products, order_product)
	}
	if err := tx.WithContext(ctx).Create(&order_products).Error; err != nil {
		return err
	}

	return nil
}
