package handler

import (
	"strconv"

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

func UpdateUser(c *fiber.Ctx) error {
	db_con := c.Locals("db_con").(*gorm.DB)
	claims, err := extract_claims_from_jwt(c)
	if err != nil {
		return err
	}
	auth_user_id := uint(claims["id"].(float64))
	o_id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	if int(auth_user_id) != o_id {
		return c.JSON(fiber.Map{
			"message": "Cannot update this user",
		})
	}
	req_body := new(schema.UserUpdateReq)
	if err := c.BodyParser(req_body); err != nil {
		return err
	}

	UserService := service.UserService{DBCon: db_con}
	o, err := UserService.UpdateUser(req_body, o_id)
	if err != nil {
		return err
	}

	res_body := new(schema.UserUpdateRes)
	copier.Copy(&res_body, &o)

	return c.JSON(res_body)
}
