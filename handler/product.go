package handler

import (
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

	o, err := controller.StoreProduct(req_body)
	if err != nil {
		return err
	}

	copier.Copy(&res_body, &o)

	return c.Status(201).JSON(res_body)
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
