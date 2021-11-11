package handler

import (
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/negadive/oneline/repository"
	"github.com/negadive/oneline/schema"
	ctrl "github.com/negadive/oneline/service"
	"gorm.io/gorm"
)

func StoreProduct(c *fiber.Ctx) error {
	db_con := c.Locals("db_con").(*gorm.DB)
	claims, err := extract_claims_from_jwt(c)
	if err != nil {
		return err
	}

	req_body := new(schema.ProductStoreReq)
	if err := c.BodyParser(&req_body); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(req_body); err != nil {
		return err
	}
	if auth_user_id := uint(claims["id"].(float64)); req_body.OwnerID != auth_user_id {
		return c.Status(403).JSON(fiber.Map{
			"message": "Cannot store product for this owner",
		})
	}

	ProductService := ctrl.ProductService{DBCon: db_con}
	o, err := ProductService.StoreProduct(req_body)
	if err != nil {
		return err
	}

	res_body := new(schema.ProductStoreRes)
	copier.Copy(&res_body, &o)

	return c.Status(201).JSON(&res_body)
}

func GetProduct(c *fiber.Ctx) error {
	db_con := c.Locals("db_con").(*gorm.DB)

	o_id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	ProductService := ctrl.ProductService{DBCon: db_con}
	o, err := ProductService.GetProduct(o_id)
	if err != nil {
		return err
	}

	res_body := new(schema.ProductGetOneRes)
	copier.Copy(&res_body, &o)

	return c.Status(200).JSON(&res_body)
}

func ListProducts(c *fiber.Ctx) error {
	db_con := c.Locals("db_con").(*gorm.DB)

	ProductService := ctrl.ProductService{DBCon: db_con}
	o, err := ProductService.ListProducts()
	if err != nil {
		return err
	}

	res_body := new([]schema.ProductListRes)
	copier.Copy(&res_body, &o)

	return c.JSON(&res_body)
}

func ListUserProducts(c *fiber.Ctx) error {
	db_con := c.Locals("db_con").(*gorm.DB)

	owner_id, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		return err
	}

	ProductService := ctrl.ProductService{DBCon: db_con}
	o, err := ProductService.ListUserProducts(owner_id)
	if err != nil {
		return err
	}

	res_body := new([]schema.ProductListRes)
	copier.Copy(&res_body, &o)

	return c.JSON(&res_body)
}

func UpdateProduct(c *fiber.Ctx) error {
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

	req_body := new(schema.ProductUpdateReq)
	if err := c.BodyParser(&req_body); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(req_body); err != nil {
		return err
	}
	ProductRepository := repository.ProductRespository{DBCon: db_con}
	if !ProductRepository.ProductWithOwnerExists(o_id, auth_user_id) {
		return c.Status(404).JSON(fiber.Map{
			"message": "No product found for this owner",
		})
	}

	ProductService := ctrl.ProductService{DBCon: db_con}
	o, err := ProductService.UpdateProduct(req_body, o_id)
	if err != nil {
		return err
	}

	res_body := new(schema.ProductUpdateRes)
	copier.Copy(&res_body, &o)

	return c.JSON(&res_body)
}

func DeleteProduct(c *fiber.Ctx) error {
	db_con := c.Locals("db_con").(*gorm.DB)
	o_id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	ProductService := ctrl.ProductService{DBCon: db_con}
	if err := ProductService.DeleteProduct(o_id); err != nil {
		return err
	}

	return c.SendStatus(204)
}
