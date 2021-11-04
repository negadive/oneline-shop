package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/negadive/oneline/controller"
	"github.com/negadive/oneline/schema"
)

func Register(c *fiber.Ctx) error {
	_o := new(schema.UserRegisterReq)

	if err := c.BodyParser(_o); err != nil {
		return err
	}

	o, err := controller.Register(_o)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{
		"name": o.Name,
	})
}
