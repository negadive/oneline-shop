package route

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/negadive/oneline/authorizer"
	"github.com/negadive/oneline/handler"
	"github.com/negadive/oneline/repository"
	"github.com/negadive/oneline/service"
	"gorm.io/gorm"
)

func setupOrderHandler(db *gorm.DB, validate *validator.Validate) handler.IOrderHandler {
	orderRepo := repository.NewOrderRepository()
	productRepo := repository.NewProductRepository()
	orderAuthzer := authorizer.NewOrderAuthorizer(orderRepo)
	orderService := service.NewOrderService(db, orderAuthzer, orderRepo, productRepo)
	orderHandler := handler.NewOrderHandler(
		orderService,
		validate,
	)

	return orderHandler
}

func Order(app *fiber.App, orderHandler handler.IOrderHandler) {
	order := app.Group("/orders")

	order.Post("/", orderHandler.Store)
}
