package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/negadive/oneline/controller"
	"github.com/negadive/oneline/schema"
)

func Login(c *fiber.Ctx) error {
	var reqBody schema.LoginReq

	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(402).JSON(fiber.Map{
			"message": "data error",
		})
	}

	token, err := controller.Login(&reqBody)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "login error",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
