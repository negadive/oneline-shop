package authorizer

import (
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/repository"
)

type IOrderAuthorizer interface {
	CanCreate(order *model.Order, user_id *uint) bool
	CanEdit(order *model.Order, user_id *uint) bool
	CanDelete(order *model.Order, user_id *uint) bool
}

type OrderAuthorizer struct {
	orderRepo repository.IOrderRepository
}

func NewOrderAuthorizer(orderRepo repository.IOrderRepository) IOrderAuthorizer {
	return &OrderAuthorizer{orderRepo: orderRepo}
}

func (oa *OrderAuthorizer) CanCreate(order *model.Order, user_id *uint) bool {
	return order.CustomerID == *user_id
}

func (oa *OrderAuthorizer) CanEdit(order *model.Order, user_id *uint) bool {
	return order.CustomerID == *user_id
}

func (oa *OrderAuthorizer) CanDelete(order *model.Order, user_id *uint) bool {
	return order.CustomerID == *user_id
}
