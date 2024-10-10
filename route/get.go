package route

import (
	"github.com/dingdinglz/dingbot/appconfig"
	"github.com/gofiber/fiber/v2"
)

func GetBotConfigRoute(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"username": appconfig.BotConfigVar.Username, "password": appconfig.BotConfigVar.Password})
}
