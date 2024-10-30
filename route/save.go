package route

import (
	"os"
	"path/filepath"

	"github.com/dingdinglz/dingbot/appconfig"
	"github.com/gofiber/fiber/v2"
)

func SaveBotConfigRoute(c *fiber.Ctx) error {
	u := c.FormValue("username")
	p := c.FormValue("password")
	appconfig.SaveBotConfig(u, p)
	return JsonMessage(c, 0, "ok")
}

func SavePluginSourceRoute(c *fiber.Ctx) error {
	rootPath, _ := os.Getwd()
	os.WriteFile(filepath.Join(rootPath, "data", "plugin", "source", c.FormValue("name")+".json"), []byte(c.FormValue("data")), os.ModePerm)
	return JsonMessage(c, 0, "ok")
}

func SavePluginRoute(c *fiber.Ctx) error {
	rootPath, _ := os.Getwd()
	os.WriteFile(filepath.Join(rootPath, "data", "plugin", c.FormValue("name")+".js"), []byte(c.FormValue("data")), os.ModePerm)
	return JsonMessage(c, 0, "ok")
}
