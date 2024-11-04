package route

import (
	"encoding/base64"
	"os"
	"path/filepath"

	"github.com/dingdinglz/dingbot/bot"
	"github.com/dingdinglz/dingbot/tool"
	"github.com/gofiber/fiber/v2"
)

func BotLoginRoute(c *fiber.Ctx) error {
	bot.PrepareBot()
	return JsonMessage(c, 0, "ok")
}

func BotFetchQRcode(c *fiber.Ctx) error {
	rootPath, _ := os.Getwd()
	for bot.DingQQBot == nil {

	}
	for !tool.FileIsExsits(filepath.Join(rootPath, "data", "qrcode.png")) && !bot.DingQQBot.Online.Load() {

	}

	if bot.DingQQBot.Online.Load() {
		return c.SendString("login success")
	}
	res, _ := os.ReadFile(filepath.Join(rootPath, "data", "qrcode.png"))
	b := base64.StdEncoding.EncodeToString(res)
	for b == "" {
		res, _ = os.ReadFile(filepath.Join(rootPath, "data", "qrcode.png"))
		b = base64.StdEncoding.EncodeToString(res)
	}
	return c.SendString("data:image/png;base64," + b)
}

func BotIsLogin(c *fiber.Ctx) error {
	if bot.DingQQBot == nil {
		return JsonMessage(c, 0, "no")
	}
	if bot.DingQQBot.Online.Load() {
		return JsonMessage(c, 0, "yes")
	}
	return JsonMessage(c, 0, "no")
}
