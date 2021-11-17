package repository

import (
	"github.com/negadive/oneline/model"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DBCon *gorm.DB
}

func (r *OrderRepository) Store(order *model.Order) error {
	result := r.DBCon.Create(order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *OrderRepository) StoreOrderProducts(order_id *uint, products *[]model.Product) error {
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
	if err := r.DBCon.Create(&order_products).Error; err != nil {
		return err
	}

	return nil
}
