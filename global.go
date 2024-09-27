package main

import "github.com/gofiber/fiber/v2"

var (
	MainServer *fiber.App
	Init       bool
	appConfig  AppConfig
	Version    string = "v2 beta"
)

type AppConfig struct {
	Port int `json:"port"`
}
