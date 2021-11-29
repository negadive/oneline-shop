package handler

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/schema"
	"github.com/negadive/oneline/service"
)

type IUserHandler interface {
	Register(fCtx *fiber.Ctx) error
	Update(fCtx *fiber.Ctx) error
}

type UserHandler struct {
	userService service.IUserService
	validate    *validator.Validate
}

func NewUserHandler(
	userService service.IUserService,
	validate *validator.Validate,
) IUserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Register(fCtx *fiber.Ctx) error {
	reqBody := new(schema.UserRegisterReq)
	if err := fCtx.BodyParser(reqBody); err != nil {
		return err
	}
	if err := h.validate.Struct(reqBody); err != nil {
		return err
	}

	user := new(model.User)
	copier.Copy(&user, &reqBody)
	err := h.userService.Register(fCtx.Context(), user)
	if err != nil {
		return err
	}

	resBody := new(schema.UserRegisterRes)
	copier.Copy(&resBody, &user)

	return fCtx.Status(201).JSON(&resBody)
}

func (h *UserHandler) Update(fCtx *fiber.Ctx) error {
	claims, err := extractClaimsFromJwt(fCtx)
	if err != nil {
		return err
	}
	authUserId := uint(claims["id"].(float64))
	userId, err := fCtx.ParamsInt("id")
	if err != nil {
		return err
	}
	uintUserId := uint(userId)

	reqBody := new(schema.UserUpdateReq)
	if err := fCtx.BodyParser(reqBody); err != nil {
		return err
	}

	user := new(model.User)
	copier.Copy(&user, &reqBody)
	err = h.userService.Update(fCtx.Context(), &authUserId, &uintUserId, user)
	if err != nil {
		return err
	}

	resBody := new(schema.UserUpdateRes)
	copier.Copy(&resBody, &user)

	return fCtx.JSON(resBody)
}
