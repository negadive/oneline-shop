package handler

import (
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/negadive/oneline/controller"
	"github.com/negadive/oneline/schema"
)

func StoreProduct(c *fiber.Ctx) error {
	req_body := new(schema.ProductStoreReq)
	res_body := new(schema.ProductStoreRes)

	if err := c.BodyParser(&req_body); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(req_body); err != nil {
		return err
	}

	o, err := controller.StoreProduct(req_body)
	if err != nil {
		return err
	}

	copier.Copy(&res_body, &o)

	return c.Status(201).JSON(&res_body)
}

func GetProduct(c *fiber.Ctx) error {
	res_body := new(schema.ProductGetOneRes)

	o_id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	o, err := controller.GetProduct(o_id)
	if err != nil {
		return err
	}

	copier.Copy(&res_body, &o)

	return c.Status(200).JSON(&res_body)
}

func ListProducts(c *fiber.Ctx) error {
	res_body := new([]schema.ProductListRes)

	o, err := controller.ListProducts()
	if err != nil {
		return err
	}

	copier.Copy(&res_body, &o)

	return c.JSON(&res_body)
}

func ListUserProducts(c *fiber.Ctx) error {
	res_body := new([]schema.ProductListRes)
	owner_id, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		return err
	}

	o, err := controller.ListUserProducts(owner_id)
	if err != nil {
		return err
	}

	copier.Copy(&res_body, &o)

	return c.JSON(&res_body)
}

func UpdateProduct(c *fiber.Ctx) error {
	o_id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	req_body := new(schema.ProductUpdateReq)
	res_body := new(schema.ProductUpdateRes)

	if err := c.BodyParser(&req_body); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(req_body); err != nil {
		return err
	}

	o, err := controller.UpdateProduct(req_body, o_id)
	if err != nil {
		return err
	}

	copier.Copy(&res_body, &o)

	return c.JSON(&res_body)
}

func DeleteProduct(c *fiber.Ctx) error {
	o_id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	if err := controller.DeleteProduct(o_id); err != nil {
		return err
	}

	return c.SendStatus(204)
}
