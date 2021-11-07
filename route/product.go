package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/negadive/oneline/handler"
)

func Product(app *fiber.App) {
	product := app.Group("/products")

	product.Post("/", handler.StoreProduct)
	product.Get("/", handler.ListProducts)
}
