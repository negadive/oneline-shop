package service

import (
	"context"
	"errors"

	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/repository"
)

type IOrderService interface {
	Store(ctx context.Context, order *model.Order, product_ids *[]uint) error
}

type OrderService struct {
	OrderRepo   repository.IOrderRepository
	ProductRepo repository.IProductRepository
}

func NewOrderService(
	order_repo repository.IOrderRepository,
	product_repo repository.IProductRepository,
) IOrderService {
	return &OrderService{
		OrderRepo:   order_repo,
		ProductRepo: product_repo,
	}
}

func (s *OrderService) Store(ctx context.Context, order *model.Order, product_ids *[]uint) error {
	products := []model.Product{}
	if err := s.ProductRepo.FindByIds(ctx, &products, product_ids); err != nil {
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
	if err := s.OrderRepo.Store(ctx, order); err != nil {
		return err
	}
	if err := s.OrderRepo.StoreOrderProducts(ctx, &order.ID, &products); err != nil {
		return err
	}

	return nil
}
