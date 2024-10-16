package route

import (
	"github.com/dingdinglz/dingbot/appconfig"
	"github.com/gofiber/fiber/v2"
)

func GenerateRenderMap(pageName string) fiber.Map {
	return fiber.Map{"Version": appconfig.Version, "Page": pageName}
}
