package route

import (
	"github.com/dingdinglz/dingbot/appconfig"
	"github.com/gofiber/fiber/v2"
)

func SaveBotConfigRoute(c *fiber.Ctx) error {
	u := c.FormValue("username")
	p := c.FormValue("password")
	appconfig.SaveBotConfig(u, p)
	return JsonMessage(c, 0, "ok")
}
