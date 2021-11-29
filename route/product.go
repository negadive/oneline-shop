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

func setupProductHandler(db *gorm.DB, validate *validator.Validate) handler.IProductHandler {
	productRepo := repository.NewProductRepository()
	productAuthzer := authorizer.NewProductAuthorizer(productRepo)
	productService := service.NewProductService(db, productAuthzer, productRepo)
	productHandler := handler.NewProductHandler(productService, validate)

	return productHandler
}

func Product(app *fiber.App, productHandler handler.IProductHandler) {
	product := app.Group("/products")

	product.Post("/", productHandler.Store)
	product.Get("/", productHandler.FindAll)
	product.Get("/:id", productHandler.GetOne)
	product.Patch("/:id", productHandler.Update)
	product.Delete("/:id", productHandler.Delete)

	users_product := app.Group("/users/:user_id/products")

	users_product.Get("/", productHandler.FindAllByUser)
}
