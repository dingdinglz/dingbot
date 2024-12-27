package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/dingdinglz/dingbot/appconfig"
	"github.com/dingdinglz/dingbot/database"
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
	tEngine.AddFunc("Open", func(Uin uint32, Typename string) bool {
		return database.OpenHave(Uin, Typename)
	})
	appconfig.MainServer = fiber.New(fiber.Config{Views: tEngine})
	appconfig.MainServer.Use(recover.New(), logger.New())
	appconfig.MainServer.Static("/", filepath.Join(rootPath, "web", "public"))
}

func ServerRun() {
	if appconfig.Init {
		ServerInitRun()
	} else {
		database.DatabaseInit()
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
	appconfig.MainServer.Use(route.LoginMiddleware)
	appconfig.MainServer.Get("/login", route.LoginRoute)
	appconfig.MainServer.Get("/", route.IndexRoute)
	appconfig.MainServer.Get("/key", route.KeyWordRoute)
	appconfig.MainServer.Get("/plugin-edit/:name", route.PluginEditRoute)
	appconfig.MainServer.Get("/group", route.GroupOpenRoute)
	appconfig.MainServer.Get("/private", route.PrivateOpenRoute)
	appconfig.MainServer.Get("/plugin", route.PluginRoute)
	appconfig.MainServer.Get("/github", route.GithubRoute)
	appconfig.MainServer.Post("/github/webhook", route.GithubWebhookRoute)

	apiRoute := appconfig.MainServer.Group("/api")
	apiRoute.Post("/login", route.ApiLoginRoute)

	apiGetRoute := apiRoute.Group("/get")
	apiGetRoute.Get("/bot", route.GetBotConfigRoute)
	apiGetRoute.Get("/plugin_source", route.GetPluginSourceRoute)

	apiSaveRoute := apiRoute.Group("/save")
	apiSaveRoute.Post("/bot", route.SaveBotConfigRoute)
	apiSaveRoute.Post("/plugin_source", route.SavePluginSourceRoute)
	apiSaveRoute.Post("/plugin", route.SavePluginRoute)

	apiBotRoute := apiRoute.Group("/bot")
	apiBotRoute.Get("/login", route.BotLoginRoute)
	apiBotRoute.Get("/qrcode", route.BotFetchQRcode)
	apiBotRoute.Get("/islogin", route.BotIsLogin)

	apiAddRoute := apiRoute.Group("/add")
	apiAddRoute.Post("/keyword", route.AddKeywordRoute)
	apiAddRoute.Post("/open", route.AddOpenRoute)
	apiAddRoute.Post("/github", route.AddGithubRoute)

	apiDeleteRoute := apiRoute.Group("/delete")
	apiDeleteRoute.Post("/keyword", route.DeleteKeywordRoute)
	apiDeleteRoute.Post("/open", route.DeleteOpenRoute)
	apiDeleteRoute.Get("/plugin_source", route.DeletePluginSourceRoute)
	apiDeleteRoute.Get("/plugin", route.DeletePluginRoute)
	apiDeleteRoute.Get("/sig", route.DeleteSigRoute)
	apiDeleteRoute.Post("/github", route.DeleteGithubRoute)

	err := appconfig.MainServer.Listen("0.0.0.0:" + strconv.Itoa(appconfig.AppConfigVar.Port))
	if err != nil {
		tool.ErrorOut("startServer", "Listen", err)
	}
}
