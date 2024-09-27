package route

import (
	"github.com/gofiber/fiber/v2"
)

func JsonMessage(c *fiber.Ctx, code int, message string) error {
	return c.JSON(fiber.Map{"code": code, "message": message})
}
