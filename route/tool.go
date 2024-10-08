package route

import (
	"github.com/dingdinglz/dingbot/appconfig"
	"github.com/gofiber/fiber/v2"
)

func GenerateRenderMap() fiber.Map {
	return fiber.Map{"Version": appconfig.Version}
}
