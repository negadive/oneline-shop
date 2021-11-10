package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func AddJWTMiddleware(app *fiber.App) {
	app.Use(jwtware.New(jwtware.Config{
		Filter: func(c *fiber.Ctx) bool {
			var is_login_req bool = (c.Path() == "/login" && c.Method() == "POST")
			var is_register_req bool = (c.Path() == "/users" && c.Method() == "POST")

			return (is_login_req || is_register_req)
		},
		SigningKey: []byte("SECRET"),
	}))
}
