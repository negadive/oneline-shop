package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/negadive/oneline/controller"
	"github.com/negadive/oneline/schema"
)

func Register(c *fiber.Ctx) error {
	req_body := new(schema.UserRegisterReq)
	res_body := new(schema.UserRegisterRes)

	if err := c.BodyParser(req_body); err != nil {
		return err
	}

	o, err := controller.Register(req_body)
	if err != nil {
		return err
	}

	copier.Copy(&res_body, &o)

	return c.Status(201).JSON(&res_body)
}
