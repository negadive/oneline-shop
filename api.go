package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/negadive/oneline/db"
	"github.com/negadive/oneline/handler"
	"github.com/negadive/oneline/middleware"
	"github.com/negadive/oneline/route"
)

func main() {
	Migrate()
	_db := db.GetDb()

	app := fiber.New(fiber.Config{
		ErrorHandler: handler.Error,
	})

	middleware.AddJWTMiddleware(app)
	route.Init(app, _db)

	app.Listen(":3000")
}
