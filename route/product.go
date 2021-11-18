package route

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/negadive/oneline/handler"
	"github.com/negadive/oneline/repository"
	"github.com/negadive/oneline/service"
	"gorm.io/gorm"
)

func Product(app *fiber.App, db *gorm.DB, validate *validator.Validate) {

	ProductRepo := repository.NewProductRepository(db)
	ProductService := service.NewProductService(ProductRepo)
	ProductHandler := handler.NewProductHandler(
		ProductService,
		validate,
	)
	product := app.Group("/products")

	product.Post("/", ProductHandler.Store)
	product.Get("/", ProductHandler.FindAll)
	product.Get("/:id", ProductHandler.GetOne)
	product.Patch("/:id", ProductHandler.Update)
	product.Delete("/:id", ProductHandler.Delete)

	users_product := app.Group("/users/:user_id/products")

	users_product.Get("/", ProductHandler.FindAllByUser)
}
