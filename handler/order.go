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
	Store(f_ctx *fiber.Ctx) error
}

type OrderHandler struct {
	OrderService service.IOrderService
	Validate     *validator.Validate
}

func NewOrderHandler(
	order_service service.IOrderService,
	validate *validator.Validate,
) IOrderHandler {
	return &OrderHandler{
		OrderService: order_service,
	}
}

func (h *OrderHandler) Store(f_ctx *fiber.Ctx) error {
	claims, err := extract_claims_from_jwt(f_ctx)
	if err != nil {
		return err
	}
	req_body := new(schema.OrderStoreReq)
	if err := f_ctx.BodyParser(&req_body); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(req_body); err != nil {
		return err
	}

	order := new(model.Order)
	copier.Copy(&order, &req_body)
	order.CustomerID = uint(claims["id"].(float64))
	if err := h.OrderService.Store(f_ctx.Context(), order, &req_body.ProductIDs); err != nil {
		return err
	}

	res_body := new(schema.OrderStoreRes)
	copier.Copy(&res_body, &order)

	return f_ctx.Status(201).JSON(&res_body)
}
