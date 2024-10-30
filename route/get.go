package route

import (
	"os"
	"path/filepath"

	"github.com/dingdinglz/dingbot/appconfig"
	"github.com/dingdinglz/dingbot/tool"
	"github.com/gofiber/fiber/v2"
)

func GetBotConfigRoute(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"username": appconfig.BotConfigVar.Username, "password": appconfig.BotConfigVar.Password})
}

func GetPluginSourceRoute(c *fiber.Ctx) error {
	rootPath, _ := os.Getwd()
	if !tool.FileIsExsits(filepath.Join(rootPath, "data", "plugin", "source", c.Query("name")+".json")) {
		return c.JSON(fiber.Map{"data": "{}"})
	}
	data, _ := os.ReadFile(filepath.Join(rootPath, "data", "plugin", "source", c.Query("name")+".json"))
	return c.JSON(fiber.Map{"data": string(data)})
}
