package handler

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/negadive/oneline/schema"
	"github.com/negadive/oneline/service"
	"gorm.io/gorm"
)

func extract_claims_from_jwt(c *fiber.Ctx) (jwt.MapClaims, error) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return claims, nil
}

func Login(c *fiber.Ctx) error {
	db_con := c.Locals("db_con").(*gorm.DB)

	reqBody := new(schema.LoginReq)
	if err := c.BodyParser(reqBody); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "data error",
		})
	}
	validate := validator.New()
	if err := validate.Struct(reqBody); err != nil {
		return err
	}

	AuthService := service.AuthService{DBCon: db_con}
	token, err := AuthService.Login(reqBody)
	if err != nil {
		return c.Status(422).JSON(fiber.Map{
			"message": "login error",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
