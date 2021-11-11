package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/negadive/oneline/db"
)

func DbCon(c *fiber.Ctx) error {
	_db := db.GetDb()

	c.Locals("db_con", _db)

	return c.Next()
}
