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

type IUserHandler interface {
	Register(f_ctx *fiber.Ctx) error
	Update(f_ctx *fiber.Ctx) error
}

type UserHandler struct {
	UserService service.IUserService
	Validate    *validator.Validate
}

func NewUserHandler(
	user_service service.IUserService,
	validate *validator.Validate,
) IUserHandler {
	return &UserHandler{
		UserService: user_service,
	}
}

func (h *UserHandler) Register(f_ctx *fiber.Ctx) error {
	req_body := new(schema.UserRegisterReq)
	if err := f_ctx.BodyParser(req_body); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(req_body); err != nil {
		return err
	}

	user := new(model.User)
	copier.Copy(&user, &req_body)
	err := h.UserService.Register(f_ctx.Context(), user)
	if err != nil {
		return err
	}

	res_body := new(schema.UserRegisterRes)
	copier.Copy(&res_body, &user)

	return f_ctx.Status(201).JSON(&res_body)
}

func (h *UserHandler) Update(f_ctx *fiber.Ctx) error {
	claims, err := extract_claims_from_jwt(f_ctx)
	if err != nil {
		return err
	}
	auth_user_id := uint(claims["id"].(float64))
	user_id, err := strconv.Atoi(f_ctx.Params("id"))
	if err != nil {
		return err
	}
	uint_user_id := uint(user_id)

	if int(auth_user_id) != user_id {
		return f_ctx.JSON(fiber.Map{
			"message": "Cannot update this user",
		})
	}
	req_body := new(schema.UserUpdateReq)
	if err := f_ctx.BodyParser(req_body); err != nil {
		return err
	}

	user := new(model.User)
	copier.Copy(&user, &req_body)
	err = h.UserService.Update(f_ctx.Context(), &uint_user_id, user)
	if err != nil {
		return err
	}

	res_body := new(schema.UserUpdateRes)
	copier.Copy(&res_body, &user)

	return f_ctx.JSON(res_body)
}
