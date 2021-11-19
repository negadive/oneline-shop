package route

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Init(app *fiber.App, db *gorm.DB) {
	validate := validator.New()

	User(app, setupUserHandler(db, validate))
	Auth(app)
	Product(app, setupProductHandler(db, validate))
	Order(app, setupOrderHandler(db, validate))
}
