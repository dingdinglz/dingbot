package route

import "github.com/gofiber/fiber/v2"

func IndexRoute(c *fiber.Ctx) error {
	pageMap := GenerateRenderMap()
	pageMap["Page"] = "bot"
	return c.Render("bot", pageMap, "layout")
}
