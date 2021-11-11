package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/negadive/oneline/handler"
)

func Product(app *fiber.App) {
	product := app.Group("/products", handler.DbCon)

	product.Post("/", handler.StoreProduct)
	product.Get("/", handler.ListProducts)
	product.Get("/:id", handler.GetProduct)
	product.Patch("/:id", handler.UpdateProduct)
	product.Delete("/:id", handler.DeleteProduct)

	users_product := app.Group("/users/:user_id/products", handler.DbCon)

	users_product.Get("/", handler.ListUserProducts)
}
