package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/negadive/oneline/handler"
)

func User(app *fiber.App) {
	user := app.Group("/users")

	user.Post("/", handler.DbCon, handler.Register)

}
