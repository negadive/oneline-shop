package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/negadive/oneline/handler"
)

func Product(app *fiber.App) {
	product := app.Group("/products")

	product.Post("/", handler.DbCon, handler.StoreProduct)
	product.Get("/", handler.DbCon, handler.ListProducts)
	product.Get("/:id", handler.DbCon, handler.GetProduct)
	product.Patch("/:id", handler.DbCon, handler.UpdateProduct)
	product.Delete("/:id", handler.DbCon, handler.DeleteProduct)

	users_product := app.Group("/users/:user_id/products")

	users_product.Get("/", handler.DbCon, handler.ListUserProducts)
}
