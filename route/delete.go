package route

import (
	"github.com/dingdinglz/dingbot/database"
	"github.com/gofiber/fiber/v2"
)

func DeleteKeywordRoute(c *fiber.Ctx) error {
	if c.FormValue("key", "") == "" {
		return JsonMessage(c, -1, "none")
	}
	database.KeyWordDelete(c.FormValue("key"))
	return JsonMessage(c, 0, "ok")
}
