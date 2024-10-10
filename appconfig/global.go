package appconfig

import (
	"github.com/gofiber/fiber/v2"
)

var (
	MainServer   *fiber.App
	Init         bool
	AppConfigVar AppConfig
	Version      string = "v2 beta"
	BotConfigVar BotConfig
)

type AppConfig struct {
	Port int `json:"port"`
}

type BotConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
