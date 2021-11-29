package handler

import (
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/negadive/oneline/custom_errors"
)

func Error(c *fiber.Ctx, err error) error {
	if nf_e, ok := err.(*custom_errors.NotFoundError); ok {
		return c.Status(404).JSON(fiber.Map{
			"message":  "resource not found",
			"resource": nf_e.Resource,
		})
	} else if f_e, ok := err.(*custom_errors.ForbiddenUser); ok {
		return c.Status(403).JSON(fiber.Map{
			"message": f_e.Error(),
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
	} else if err_text := err.Error(); err_text == "Unprocessable Entity" {
		return c.Status(400).JSON(fiber.Map{
			"message": err_text,
		})
	}

	return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
}
