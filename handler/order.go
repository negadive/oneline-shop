package handler

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/repository"
	"github.com/negadive/oneline/schema"
	"github.com/negadive/oneline/service"
	"gorm.io/gorm"
)

func StoreOrder(c *fiber.Ctx) error {
	db_con := c.Locals("db_con").(*gorm.DB)
	claims, err := extract_claims_from_jwt(c)
	if err != nil {
		return err
	}
	order_repo := repository.OrderRepository{DBCon: db_con}
	product_repo := repository.ProductRespository{DBCon: db_con}

	req_body := new(schema.OrderStoreReq)
	if err := c.BodyParser(&req_body); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(req_body); err != nil {
		return err
	}

	order := new(model.Order)
	copier.Copy(&order, &req_body)
	order.CustomerID = uint(claims["id"].(float64))
	OrderService := service.OrderService{
		DBCon:       db_con,
		OrderRepo:   &order_repo,
		ProductRepo: &product_repo,
	}
	if err := OrderService.Store(order, &req_body.ProductIDs); err != nil {
		return err
	}

	res_body := new(schema.OrderStoreRes)
	copier.Copy(&res_body, &order)

	return c.Status(201).JSON(&res_body)
}
