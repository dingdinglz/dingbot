package route

import (
	"github.com/dingdinglz/dingbot/database"
	"github.com/dingdinglz/dingbot/tool"
	"github.com/gofiber/fiber/v2"
)

func AddKeywordRoute(c *fiber.Ctx) error {
	if c.FormValue("key", "") == "" || c.FormValue("text") == "" {
		return JsonMessage(c, -1, "none")
	}
	var p bool = true
	if c.FormValue("precise") == "false" {
		p = false
	}
	database.KeyWordCreate(c.FormValue("key"), c.FormValue("text"), p)
	return JsonMessage(c, 0, "ok")
}

func AddOpenRoute(c *fiber.Ctx) error {
	if c.FormValue("type", "") == "" || c.FormValue("uin", "") == "" {
		return JsonMessage(c, -1, "none")
	}
	database.OpenAdd(tool.StringToUint32(c.FormValue("uin", "")), c.FormValue("type", ""))
	return JsonMessage(c, 0, "ok")
}
