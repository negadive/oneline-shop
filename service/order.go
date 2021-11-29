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
	Store(ctx context.Context, actor_id *uint, order *model.Order, product_ids *[]uint) error
}

type OrderService struct {
	DBCon        *gorm.DB
	orderAuthzer authorizer.IOrderAuthorizer
	OrderRepo    repository.IOrderRepository
	ProductRepo  repository.IProductRepository
}

func NewOrderService(
	db_con *gorm.DB,
	orderAuthzer authorizer.IOrderAuthorizer,
	order_repo repository.IOrderRepository,
	product_repo repository.IProductRepository,
) IOrderService {
	return &OrderService{
		DBCon:        db_con,
		orderAuthzer: orderAuthzer,
		OrderRepo:    order_repo,
		ProductRepo:  product_repo,
	}
}

func (s *OrderService) Store(ctx context.Context, actor_id *uint, order *model.Order, product_ids *[]uint) error {
	tx := s.DBCon.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	products := []model.Product{}
	if err := s.ProductRepo.FindByIds(ctx, tx, &products, product_ids); err != nil {
		return err
	}
	if len(products) != len(*product_ids) {
		tx.Rollback()
		return errors.New("some products not found")
	}
	total_price := 0.0
	for _, product := range products {
		total_price += product.Price
	}

	order.Status = "CREATED"
	order.TotalPrice = total_price
	if err := s.OrderRepo.Store(ctx, tx, order); err != nil {
		tx.Rollback()
		return err
	}
	if err := s.OrderRepo.StoreOrderProducts(ctx, tx, &order.ID, &products); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
