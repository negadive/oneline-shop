package service

import (
	"context"
	"errors"

	"github.com/negadive/oneline/authorizer"
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/repository"
	"gorm.io/gorm"
)

type IOrderService interface {
	Store(ctx context.Context, actorId *uint, order *model.Order, productIds *[]uint) error
}

type OrderService struct {
	dBCon        *gorm.DB
	orderAuthzer authorizer.IOrderAuthorizer
	orderRepo    repository.IOrderRepository
	productRepo  repository.IProductRepository
}

func NewOrderService(
	dBCon *gorm.DB,
	orderAuthzer authorizer.IOrderAuthorizer,
	orderRepo repository.IOrderRepository,
	productRepo repository.IProductRepository,
) IOrderService {
	return &OrderService{
		dBCon:        dBCon,
		orderAuthzer: orderAuthzer,
		orderRepo:    orderRepo,
		productRepo:  productRepo,
	}
}

func (s *OrderService) Store(ctx context.Context, actorId *uint, order *model.Order, productIds *[]uint) error {
	tx := s.dBCon.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	products := []model.Product{}
	if err := s.productRepo.FindByIds(ctx, tx, &products, productIds); err != nil {
		return err
	}
	if len(products) != len(*productIds) {
		tx.Rollback()
		return errors.New("some products not found")
	}
	totalPrice := 0.0
	for _, product := range products {
		totalPrice += product.Price
	}

	order.Status = "CREATED"
	order.TotalPrice = totalPrice
	if err := s.orderRepo.Store(ctx, tx, order); err != nil {
		tx.Rollback()
		return err
	}
	if err := s.orderRepo.StoreOrderProducts(ctx, tx, &order.ID, &products); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
