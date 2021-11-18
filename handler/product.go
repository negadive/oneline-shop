package handler

import (
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/schema"
	"github.com/negadive/oneline/service"
)

type IProductHandler interface {
	GetOne(f_ctx *fiber.Ctx) error
	Store(f_ctx *fiber.Ctx) error
	Update(f_ctx *fiber.Ctx) error
	Delete(f_ctx *fiber.Ctx) error
	FindAll(f_ctx *fiber.Ctx) error
	FindAllByUser(f_ctx *fiber.Ctx) error
}
type ProductHandler struct {
	ProductService service.IProductService
	Validate       *validator.Validate
}

func NewProductHandler(product_service service.IProductService, validate *validator.Validate) IProductHandler {
	h := ProductHandler{
		ProductService: product_service,
		Validate:       validate,
	}

	return &h
}

func (h *ProductHandler) Store(f_ctx *fiber.Ctx) error {
	claims, err := extract_claims_from_jwt(f_ctx)
	if err != nil {
		return err
	}

	req_body := new(schema.ProductStoreReq)
	if err := f_ctx.BodyParser(&req_body); err != nil {
		return err
	}
	if err := h.Validate.Struct(req_body); err != nil {
		return err
	}
	if auth_user_id := uint(claims["id"].(float64)); req_body.OwnerID != auth_user_id {
		return f_ctx.Status(403).JSON(fiber.Map{
			"message": "Cannot store product for this owner",
		})
	}

	product := new(model.Product)
	copier.Copy(&product, &req_body)
	if err := h.ProductService.Store(f_ctx.Context(), product); err != nil {
		return err
	}

	res_body := new(schema.ProductStoreRes)
	copier.Copy(&res_body, &product)

	return f_ctx.Status(201).JSON(&res_body)
}

func (h *ProductHandler) GetOne(f_ctx *fiber.Ctx) error {
	product_id, err := strconv.Atoi(f_ctx.Params("id"))
	if err != nil {
		return err
	}
	uint_product_id := uint(product_id)
	product, err := h.ProductService.GetOne(f_ctx.Context(), &uint_product_id)
	if err != nil {
		return err
	}

	res_body := new(schema.ProductGetOneRes)
	copier.Copy(&res_body, &product)

	return f_ctx.Status(200).JSON(&res_body)
}

func (h *ProductHandler) FindAll(f_ctx *fiber.Ctx) error {
	product, err := h.ProductService.FindAll(f_ctx.Context())
	if err != nil {
		return err
	}

	res_body := new([]schema.ProductListRes)
	copier.Copy(&res_body, &product)

	return f_ctx.JSON(&res_body)
}

func (h *ProductHandler) FindAllByUser(c *fiber.Ctx) error {
	owner_id, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		return err
	}
	uint_owner_id := uint(owner_id)

	product, err := h.ProductService.FindAllByUser(c.Context(), &uint_owner_id)
	if err != nil {
		return err
	}

	res_body := new([]schema.ProductListRes)
	copier.Copy(&res_body, &product)

	return c.JSON(&res_body)
}

func (h *ProductHandler) Update(f_ctx *fiber.Ctx) error {
	claims, err := extract_claims_from_jwt(f_ctx)
	if err != nil {
		return err
	}
	auth_user_id := uint(claims["id"].(float64))
	product_id, err := strconv.Atoi(f_ctx.Params("id"))
	if err != nil {
		return err
	}
	uint_product_id := uint(product_id)

	req_body := new(schema.ProductUpdateReq)
	if err := f_ctx.BodyParser(&req_body); err != nil {
		return err
	}
	if err := h.Validate.Struct(req_body); err != nil {
		return err
	}
	if !h.ProductService.GetProductRepo().ProductWithOwnerExists(f_ctx.Context(), &uint_product_id, &auth_user_id) {
		return f_ctx.Status(404).JSON(fiber.Map{
			"message": "No product found for this owner",
		})
	}

	product := new(model.Product)
	copier.Copy(&product, &req_body)
	if err := h.ProductService.Update(f_ctx.Context(), &uint_product_id, product); err != nil {
		return err
	}

	res_body := new(schema.ProductUpdateRes)
	copier.Copy(&res_body, &product)

	return f_ctx.JSON(&res_body)
}

func (h *ProductHandler) Delete(f_ctx *fiber.Ctx) error {
	product_id, err := strconv.Atoi(f_ctx.Params("id"))
	if err != nil {
		return err
	}
	uint_product_id := uint(product_id)

	if err := h.ProductService.Delete(f_ctx.Context(), &uint_product_id); err != nil {
		return err
	}

	return f_ctx.SendStatus(204)
}
