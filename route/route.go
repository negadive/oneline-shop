package route

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Init(app *fiber.App, db *gorm.DB) {
	validate := validator.New()

	User(app)
	Auth(app)
	Product(app, db, validate)
	Order(app, db, validate)
}
