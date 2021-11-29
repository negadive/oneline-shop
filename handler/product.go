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
	GetOne(fCtx *fiber.Ctx) error
	Store(fCtx *fiber.Ctx) error
	Update(fCtx *fiber.Ctx) error
	Delete(fCtx *fiber.Ctx) error
	FindAll(fCtx *fiber.Ctx) error
	FindAllByUser(fCtx *fiber.Ctx) error
}
type ProductHandler struct {
	productService service.IProductService
	Validate       *validator.Validate
}

func NewProductHandler(productService service.IProductService, validate *validator.Validate) IProductHandler {
	return &ProductHandler{
		productService: productService,
		Validate:       validate,
	}
}

func (h *ProductHandler) Store(fCtx *fiber.Ctx) error {
	claims, err := extractClaimsFromJwt(fCtx)
	if err != nil {
		return err
	}
	authUserId := uint(claims["id"].(float64))

	reqBody := new(schema.ProductStoreReq)
	if err := fCtx.BodyParser(&reqBody); err != nil {
		return err
	}
	if err := h.Validate.Struct(reqBody); err != nil {
		return err
	}

	product := new(model.Product)
	copier.Copy(&product, &reqBody)
	if err := h.productService.Store(fCtx.Context(), &authUserId, product); err != nil {
		return err
	}

	resBody := new(schema.ProductStoreRes)
	copier.Copy(&resBody, &product)

	return fCtx.Status(201).JSON(&resBody)
}

func (h *ProductHandler) GetOne(fCtx *fiber.Ctx) error {
	productId, err := strconv.Atoi(fCtx.Params("id"))
	if err != nil {
		return err
	}
	uintProductId := uint(productId)
	product, err := h.productService.GetOne(fCtx.Context(), &uintProductId)
	if err != nil {
		return err
	}

	resBody := new(schema.ProductGetOneRes)
	copier.Copy(&resBody, &product)

	return fCtx.Status(200).JSON(&resBody)
}

func (h *ProductHandler) FindAll(fCtx *fiber.Ctx) error {
	product, err := h.productService.FindAll(fCtx.Context())
	if err != nil {
		return err
	}

	resBody := new([]schema.ProductListRes)
	copier.Copy(&resBody, &product)

	return fCtx.JSON(&resBody)
}

func (h *ProductHandler) FindAllByUser(c *fiber.Ctx) error {
	ownerId, err := c.ParamsInt("user_id")
	if err != nil {
		return err
	}
	uintOwnerId := uint(ownerId)

	product, err := h.productService.FindAllByUser(c.Context(), &uintOwnerId)
	if err != nil {
		return err
	}

	resBody := new([]schema.ProductListRes)
	copier.Copy(&resBody, &product)

	return c.JSON(&resBody)
}

func (h *ProductHandler) Update(fCtx *fiber.Ctx) error {
	claims, err := extractClaimsFromJwt(fCtx)
	if err != nil {
		return err
	}
	authUserId := uint(claims["id"].(float64))

	productId, err := fCtx.ParamsInt("id")
	if err != nil {
		return err
	}
	uintProductId := uint(productId)

	reqBody := new(schema.ProductUpdateReq)
	if err := fCtx.BodyParser(&reqBody); err != nil {
		return err
	}
	if err := h.Validate.Struct(reqBody); err != nil {
		return err
	}

	product := new(model.Product)
	copier.Copy(&product, &reqBody)
	if err := h.productService.Update(fCtx.Context(), &authUserId, &uintProductId, product); err != nil {
		return err
	}

	resBody := new(schema.ProductUpdateRes)
	copier.Copy(&resBody, &product)

	return fCtx.JSON(&resBody)
}

func (h *ProductHandler) Delete(fCtx *fiber.Ctx) error {
	claims, err := extractClaimsFromJwt(fCtx)
	if err != nil {
		return err
	}
	authUserId := uint(claims["id"].(float64))

	productId, err := fCtx.ParamsInt("id")
	if err != nil {
		return err
	}
	uintProductId := uint(productId)

	if err := h.productService.Delete(fCtx.Context(), &authUserId, &uintProductId); err != nil {
		return err
	}

	return fCtx.SendStatus(204)
}
