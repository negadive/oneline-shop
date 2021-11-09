package handler

import (
	"errors"
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Error(c *fiber.Ctx, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"message": "resource not found",
		})
	} else if strings.Contains(err.Error(), "duplicate key") {
		return c.Status(409).JSON(fiber.Map{
			"message": "duplicate resource",
		})
	} else if v_err, ok := err.(validator.ValidationErrors); ok {
		var details []map[string]interface{}
		for _, v_err2 := range v_err {
			detail := map[string]interface{}{
				"field": v_err2.StructNamespace(),
				"tag":   v_err2.Tag(),
				"value": v_err2.Param(),
			}
			details = append(details, detail)
		}

		return c.Status(422).JSON(fiber.Map{
			"message": "invalid data",
			"details": details,
		})
	}

	return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
}
