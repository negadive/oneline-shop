package handler

import (
	"errors"
	"strings"

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
			"message": "duplicare resource",
		})
	}

	return nil
}
