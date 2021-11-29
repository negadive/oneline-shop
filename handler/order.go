package handler

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/schema"
	"github.com/negadive/oneline/service"
)

type IOrderHandler interface {
	Store(fCtx *fiber.Ctx) error
}

type OrderHandler struct {
	orderService service.IOrderService
	validate     *validator.Validate
}

func NewOrderHandler(
	orderService service.IOrderService,
	validate *validator.Validate,
) IOrderHandler {
	return &OrderHandler{
		orderService: orderService,
		validate:     validate,
	}
}

func (h *OrderHandler) Store(fCtx *fiber.Ctx) error {
	claims, err := extractClaimsFromJwt(fCtx)
	if err != nil {
		return err
	}
	authUserId := uint(claims["id"].(float64))

	reqBody := new(schema.OrderStoreReq)
	if err := fCtx.BodyParser(&reqBody); err != nil {
		return err
	}
	if err := h.validate.Struct(reqBody); err != nil {
		return err
	}

	order := new(model.Order)
	copier.Copy(&order, &reqBody)
	order.CustomerID = authUserId
	if err := h.orderService.Store(fCtx.Context(), &authUserId, order, &reqBody.ProductIDs); err != nil {
		return err
	}

	resBody := new(schema.OrderStoreRes)
	copier.Copy(&resBody, &order)

	return fCtx.Status(201).JSON(&resBody)
}
