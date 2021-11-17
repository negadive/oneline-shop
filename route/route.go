package route

import (
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	User(app)
	Auth(app)
	Product(app)
	Order(app)
}
