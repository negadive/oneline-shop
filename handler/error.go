package handler

import (
	"strings"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/negadive/oneline/customErrors"
)

func Error(c *fiber.Ctx, err error) error {
	if nfE, ok := err.(*customErrors.NotFoundError); ok {
		return c.Status(404).JSON(fiber.Map{
			"message":  "resource not found",
			"resource": nfE.Resource,
		})
	} else if fE, ok := err.(*customErrors.ForbiddenUser); ok {
		return c.Status(403).JSON(fiber.Map{
			"message": fE.Error(),
		})
	} else if strings.Contains(err.Error(), "duplicate key") {
		return c.Status(409).JSON(fiber.Map{
			"message": "duplicate resource",
		})
	} else if vErr, ok := err.(validator.ValidationErrors); ok {
		var details []map[string]interface{}
		for _, vErr2 := range vErr {
			detail := map[string]interface{}{
				"field": vErr2.StructNamespace(),
				"tag":   vErr2.Tag(),
				"value": vErr2.Param(),
			}
			details = append(details, detail)
		}

		return c.Status(422).JSON(fiber.Map{
			"message": "invalid data",
			"details": details,
		})
	} else if errText := err.Error(); errText == "Unprocessable Entity" {
		return c.Status(400).JSON(fiber.Map{
			"message": errText,
		})
	}

	return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
}
