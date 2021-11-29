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
	OrderRepo := repository.NewOrderRepository()
	ProductRepo := repository.NewProductRepository()
	OrderAuthzer := authorizer.NewOrderAuthorizer(OrderRepo)
	OrderService := service.NewOrderService(db, OrderAuthzer, OrderRepo, ProductRepo)
	OrderHandler := handler.NewOrderHandler(
		OrderService,
		validate,
	)

	return OrderHandler
}

func Order(app *fiber.App, OrderHandler handler.IOrderHandler) {
	order := app.Group("/orders")

	order.Post("/", OrderHandler.Store)
}
