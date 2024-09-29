package route

import "github.com/gofiber/fiber/v2"

func GenerateRenderMap() fiber.Map {
	return fiber.Map{"Version": Version}
}
