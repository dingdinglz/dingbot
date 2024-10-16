package route

import (
	"github.com/dingdinglz/dingbot/database"
	"github.com/gofiber/fiber/v2"
)

func IndexRoute(c *fiber.Ctx) error {
	pageMap := GenerateRenderMap("bot")
	return c.Render("bot", pageMap, "layout")
}

func KeyWordRoute(c *fiber.Ctx) error {
	pageMap := GenerateRenderMap("key")
	pageMap["Keys"] = database.KeyWordGetAll()
	return c.Render("key", pageMap, "layout")
}
