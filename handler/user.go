package handler

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/negadive/oneline/schema"
	"github.com/negadive/oneline/service"
)

func Register(c *fiber.Ctx) error {
	req_body := new(schema.UserRegisterReq)
	res_body := new(schema.UserRegisterRes)

	if err := c.BodyParser(req_body); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(req_body); err != nil {
		return err
	}

	o, err := service.Register(req_body)
	if err != nil {
		return err
	}

	copier.Copy(&res_body, &o)

	return c.Status(201).JSON(&res_body)
}
