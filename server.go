package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/dingdinglz/dingbot/route"
	"github.com/dingdinglz/dingbot/tool"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func ServerInit() {
	rootPath, _ := os.Getwd()
	MainServer = fiber.New(fiber.Config{})
	MainServer.Use(recover.New(), logger.New())
	MainServer.Static("/", filepath.Join(rootPath, "web", "public"))
}

func ServerRun() {
	if Init {
		ServerInitRun()
	} else {
		ServerCommonRun()
	}
}

func ServerInitRun() {
	rootPath, _ := os.Getwd()
	fmt.Println("Init Mode:Please open http://127.0.0.1:7000/ to init the program.")
	apiRoute := MainServer.Group("/api")
	apiRoute.Post("/init", route.InitRoute)
	MainServer.Static("/", filepath.Join(rootPath, "web", "init", "index.html"))
	err := MainServer.Listen("0.0.0.0:" + strconv.Itoa(appConfig.Port))
	if err != nil {
		tool.ErrorOut("startServer", "Listen", err)
	}
}

func ServerCommonRun() {
	err := MainServer.Listen("0.0.0.0:" + strconv.Itoa(appConfig.Port))
	if err != nil {
		tool.ErrorOut("startServer", "Listen", err)
	}
}
