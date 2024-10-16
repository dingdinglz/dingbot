package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/dingdinglz/dingbot/appconfig"
	"github.com/dingdinglz/dingbot/route"
	"github.com/dingdinglz/dingbot/tool"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func ServerInit() {
	rootPath, _ := os.Getwd()
	tEngine := html.New(filepath.Join(rootPath, "web", "index"), ".html")
	tEngine.ShouldReload = true
	appconfig.MainServer = fiber.New(fiber.Config{Views: tEngine})
	appconfig.MainServer.Use(recover.New(), logger.New())
	appconfig.MainServer.Static("/", filepath.Join(rootPath, "web", "public"))
}

func ServerRun() {
	if appconfig.Init {
		ServerInitRun()
	} else {
		ServerCommonRun()
	}
}

func ServerInitRun() {
	rootPath, _ := os.Getwd()
	fmt.Println("Init Mode:Please open http://127.0.0.1:7000/ to init the program.")
	apiRoute := appconfig.MainServer.Group("/api")
	apiRoute.Post("/init", route.InitRoute)
	appconfig.MainServer.Static("/", filepath.Join(rootPath, "web", "init", "index.html"))
	err := appconfig.MainServer.Listen("0.0.0.0:" + strconv.Itoa(appconfig.AppConfigVar.Port))
	if err != nil {
		tool.ErrorOut("startServer", "Listen", err)
	}
}

func ServerCommonRun() {
	appconfig.MainServer.Get("/", route.IndexRoute)
	appconfig.MainServer.Get("/key", route.KeyWordRoute)

	apiRoute := appconfig.MainServer.Group("/api")

	apiGetRoute := apiRoute.Group("/get")
	apiGetRoute.Get("/bot", route.GetBotConfigRoute)

	apiSaveRoute := apiRoute.Group("/save")
	apiSaveRoute.Post("/bot", route.SaveBotConfigRoute)

	apiBotRoute := apiRoute.Group("/bot")
	apiBotRoute.Get("/login", route.BotLoginRoute)
	apiBotRoute.Get("/qrcode", route.BotFetchQRcode)
	apiBotRoute.Get("/islogin", route.BotIsLogin)

	apiAddRoute := apiRoute.Group("/add")
	apiAddRoute.Post("/keyword", route.AddKeywordRoute)

	apiDeleteRoute := apiRoute.Group("/delete")
	apiDeleteRoute.Post("/keyword", route.DeleteKeywordRoute)

	err := appconfig.MainServer.Listen("0.0.0.0:" + strconv.Itoa(appconfig.AppConfigVar.Port))
	if err != nil {
		tool.ErrorOut("startServer", "Listen", err)
	}
}
