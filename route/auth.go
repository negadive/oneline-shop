package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/negadive/oneline/handler"
)

func Auth(app *fiber.App) {
	app.Post("/login", handler.Login)
}
