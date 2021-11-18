package handler

import (
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/repository"
	"github.com/negadive/oneline/schema"
	"github.com/negadive/oneline/service"
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

	product := new(model.Product)
	copier.Copy(&product, &req_body)
	ProductService := service.ProductService{
		ProductRepo: &repository.ProductRespository{DBCon: db_con},
	}
	if err := ProductService.StoreProduct(product); err != nil {
		return err
	}

	res_body := new(schema.ProductStoreRes)
	copier.Copy(&res_body, &product)

	return c.Status(201).JSON(&res_body)
}

func GetProduct(c *fiber.Ctx) error {
	db_con := c.Locals("db_con").(*gorm.DB)

	product_id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	uint_product_id := uint(product_id)

	ProductService := service.ProductService{
		ProductRepo: &repository.ProductRespository{DBCon: db_con},
	}
	product, err := ProductService.GetProduct(&uint_product_id)
	if err != nil {
		return err
	}

	res_body := new(schema.ProductGetOneRes)
	copier.Copy(&res_body, &product)

	return c.Status(200).JSON(&res_body)
}

func ListProducts(c *fiber.Ctx) error {
	db_con := c.Locals("db_con").(*gorm.DB)

	ProductService := service.ProductService{
		ProductRepo: &repository.ProductRespository{DBCon: db_con},
	}
	product, err := ProductService.ListProducts()
	if err != nil {
		return err
	}

	res_body := new([]schema.ProductListRes)
	copier.Copy(&res_body, &product)

	return c.JSON(&res_body)
}

func ListUserProducts(c *fiber.Ctx) error {
	db_con := c.Locals("db_con").(*gorm.DB)

	owner_id, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		return err
	}
	uint_owner_id := uint(owner_id)

	ProductService := service.ProductService{
		ProductRepo: &repository.ProductRespository{DBCon: db_con},
	}
	product, err := ProductService.ListUserProducts(&uint_owner_id)
	if err != nil {
		return err
	}

	res_body := new([]schema.ProductListRes)
	copier.Copy(&res_body, &product)

	return c.JSON(&res_body)
}

func UpdateProduct(c *fiber.Ctx) error {
	db_con := c.Locals("db_con").(*gorm.DB)
	claims, err := extract_claims_from_jwt(c)
	if err != nil {
		return err
	}
	auth_user_id := uint(claims["id"].(float64))
	product_id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	uint_product_id := uint(product_id)

	req_body := new(schema.ProductUpdateReq)
	if err := c.BodyParser(&req_body); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(req_body); err != nil {
		return err
	}
	ProductRepository := repository.ProductRespository{DBCon: db_con}
	if !ProductRepository.ProductWithOwnerExists(&uint_product_id, &auth_user_id) {
		return c.Status(404).JSON(fiber.Map{
			"message": "No product found for this owner",
		})
	}

	product := new(model.Product)
	copier.Copy(&product, &req_body)
	ProductService := service.ProductService{
		ProductRepo: &ProductRepository,
	}
	if err := ProductService.UpdateProduct(product, &uint_product_id); err != nil {
		return err
	}

	res_body := new(schema.ProductUpdateRes)
	copier.Copy(&res_body, &product)

	return c.JSON(&res_body)
}

func DeleteProduct(c *fiber.Ctx) error {
	db_con := c.Locals("db_con").(*gorm.DB)
	product_id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}
	uint_product_id := uint(product_id)

	ProductService := service.ProductService{
		ProductRepo: &repository.ProductRespository{DBCon: db_con},
	}
	if err := ProductService.DeleteProduct(&uint_product_id); err != nil {
		return err
	}

	return c.SendStatus(204)
}
