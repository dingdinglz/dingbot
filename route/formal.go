package route

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/dingdinglz/dingbot/bot"
	"github.com/dingdinglz/dingbot/database"
	"github.com/dingdinglz/dingbot/plugin"
	"github.com/gofiber/fiber/v2"
)

func IndexRoute(c *fiber.Ctx) error {
	pageMap := GenerateRenderMap("bot")
	return c.Render("bot", pageMap, "layout")
}

func KeyWordRoute(c *fiber.Ctx) error {
	pageMap := GenerateRenderMap("key")
	pageMap["Keys"] = database.KeyWordGetAll()
	return c.Render("key", pageMap, "layout")
}

func PluginEditRoute(c *fiber.Ctx) error {
	pageMap := GenerateRenderMap("plugin-edit")
	pageMap["PluginName"] = c.Params("name")
	return c.Render("plugin-edit", pageMap, "layout")
}

type qqGroup struct {
	Name        string
	Uin         uint32
	MemberCount uint32
}

func GroupOpenRoute(c *fiber.Ctx) error {
	if bot.DingQQBot == nil {
		return c.Redirect("/")
	}
	if !bot.DingQQBot.Online.Load() {
		return c.Redirect("/")
	}
	groups, _ := bot.DingQQBot.GetAllGroupsInfo()
	var qqgs []qqGroup
	for _, i := range groups {
		var _cnt qqGroup
		_cnt.Name = i.GroupName
		_cnt.Uin = i.GroupUin
		_cnt.MemberCount = i.MemberCount
		qqgs = append(qqgs, _cnt)
	}
	pageMap := GenerateRenderMap("group")
	pageMap["GroupList"] = qqgs
	return c.Render("group", pageMap, "layout")
}

type qqFriend struct {
	Name   string
	Uin    uint32
	Avatar string
}

func PrivateOpenRoute(c *fiber.Ctx) error {
	if bot.DingQQBot == nil {
		return c.Redirect("/")
	}
	if !bot.DingQQBot.Online.Load() {
		return c.Redirect("/")
	}
	friends, _ := bot.DingQQBot.GetFriendsData()
	var qqfs []qqFriend
	for _, i := range friends {
		var _cnt qqFriend
		_cnt.Name = i.Nickname
		_cnt.Uin = i.Uin
		_cnt.Avatar = i.Avatar
		qqfs = append(qqfs, _cnt)
	}
	pageMap := GenerateRenderMap("private")
	pageMap["FriendList"] = qqfs
	return c.Render("private", pageMap, "layout")
}

func PluginRoute(c *fiber.Ctx) error {
	rootPath, _ := os.Getwd()
	var PluginSourceList []string
	filepath.Walk(filepath.Join(rootPath, "data", "plugin", "source"), func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".json" {
			PluginSourceList = append(PluginSourceList, strings.ReplaceAll(filepath.Base(path), filepath.Ext(path), ""))
		}
		return nil
	})
	pageMap := GenerateRenderMap("plugin")
	pageMap["PluginSourceList"] = PluginSourceList
	pageMap["PluginList"] = plugin.GetAllPluginsInfos()
	return c.Render("plugin", pageMap, "layout")
}

func GithubRoute(ctx *fiber.Ctx) error {
	if bot.DingQQBot == nil {
		return ctx.Redirect("/")
	}
	if !bot.DingQQBot.Online.Load() {
		return ctx.Redirect("/")
	}
	pageMap := GenerateRenderMap("webhook-github")
	return ctx.Render("github", pageMap, "layout")
}
