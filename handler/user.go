package handler

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/negadive/oneline/schema"
	"github.com/negadive/oneline/service"
	"gorm.io/gorm"
)

func Register(c *fiber.Ctx) error {
	db_con := c.Locals("db_con").(*gorm.DB)

	req_body := new(schema.UserRegisterReq)
	if err := c.BodyParser(req_body); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(req_body); err != nil {
		return err
	}

	UserService := service.UserService{DBCon: db_con}
	o, err := UserService.Register(req_body)
	if err != nil {
		return err
	}

	res_body := new(schema.UserRegisterRes)
	copier.Copy(&res_body, &o)

	return c.Status(201).JSON(&res_body)
}
