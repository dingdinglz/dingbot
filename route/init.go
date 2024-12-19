package route

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/dingdinglz/dingbot/tool"
	"github.com/gofiber/fiber/v2"
)

func InitRoute(c *fiber.Ctx) error {
	port := c.FormValue("port", "")
	databaseType := c.FormValue("database", "")
	source := c.FormValue("source", "")
	if port == "" {
		return JsonMessage(c, 1, "port不能为空!")
	}
	if databaseType == "" {
		return JsonMessage(c, 1, "database不能为空!")
	}
	if source == "" {
		return JsonMessage(c, 1, "source不能为空!")
	}
	configText, _ := json.Marshal(fiber.Map{"port": tool.StringToInt(port), "database": databaseType, "source": source})
	rootPath, _ := os.Getwd()
	os.WriteFile(filepath.Join(rootPath, "data", "config.json"), configText, os.ModePerm)
	return JsonMessage(c, 0, "ok")
}
