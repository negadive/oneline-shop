package service

import (
	"errors"

	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/repository"
	"gorm.io/gorm"
)

type OrderService struct {
	DBCon       *gorm.DB
	OrderRepo   *repository.OrderRepository
	ProductRepo *repository.ProductRespository
}

func (s *OrderService) Store(order *model.Order, product_ids *[]uint) error {
	products := []model.Product{}
	if err := s.ProductRepo.FindByIds(&products, product_ids); err != nil {
		return err
	}
	if len(products) != len(*product_ids) {
		return errors.New("some products not found")
	}
	total_price := 0.0
	for _, product := range products {
		total_price += product.Price
	}

	order.Status = "CREATED"
	order.TotalPrice = total_price
	if err := s.OrderRepo.Store(order); err != nil {
		return err
	}
	if err := s.OrderRepo.StoreOrderProducts(&order.ID, &products); err != nil {
		return err
	}

	return nil
}
