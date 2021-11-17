package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/negadive/oneline/handler"
)

func Order(app *fiber.App) {
	order := app.Group("/orders")

	order.Post("/", handler.DbCon, handler.StoreOrder)
}
