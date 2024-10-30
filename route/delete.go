package route

import (
	"github.com/dingdinglz/dingbot/database"
	"github.com/dingdinglz/dingbot/tool"
	"github.com/gofiber/fiber/v2"
)

func DeleteKeywordRoute(c *fiber.Ctx) error {
	if c.FormValue("key", "") == "" {
		return JsonMessage(c, -1, "none")
	}
	database.KeyWordDelete(c.FormValue("key"))
	return JsonMessage(c, 0, "ok")
}

func DeleteOpenRoute(c *fiber.Ctx) error {
	if c.FormValue("type", "") == "" || c.FormValue("uin", "") == "" {
		return JsonMessage(c, -1, "none")
	}
	database.OpenDelete(tool.StringToUint32(c.FormValue("uin", "")), c.FormValue("type", ""))
	return JsonMessage(c, 0, "ok")
}
