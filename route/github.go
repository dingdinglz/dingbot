package route

import (
	"encoding/json"
	"strings"

	"github.com/LagrangeDev/LagrangeGo/message"
	"github.com/dingdinglz/dingbot/bot"
	"github.com/dingdinglz/dingbot/database"
	"github.com/dingdinglz/dingbot/tool"
	"github.com/gofiber/fiber/v2"
)

func GithubWebhookRoute(ctx *fiber.Ctx) error {
	if bot.DingQQBot == nil {
		return JsonMessage(ctx, 0, "ok")
	}
	if !bot.DingQQBot.Online.Load() {
		return JsonMessage(ctx, 0, "ok")
	}
	eventName := ctx.Get("X-GitHub-Event")
	var webhookInfos map[string]interface{}
	json.Unmarshal(ctx.Body(), &webhookInfos)
	repositoryName := webhookInfos["repository"].(map[string]interface{})["full_name"].(string)
	if !database.GithubWebhookExist(repositoryName) {
		return JsonMessage(ctx, 0, "ok")
	}
	sendGroupsString := database.GithubWebhookGet(repositoryName)
	sendGroups := strings.Split(sendGroupsString, ",")
	switch eventName {
	case "push":
		bot.DingQQBot.SendGroupMessage(11, []message.IMessageElement{message.NewText("11")})
	case "star":
		if webhookInfos["action"].(string) == "created" {
			if database.StarExist(repositoryName, webhookInfos["sender"].(map[string]interface{})["login"].(string)) {
				database.StarAdd(repositoryName, webhookInfos["sender"].(map[string]interface{})["login"].(string))
				messageSend := "项目" + repositoryName + "得到" + webhookInfos["sender"].(map[string]interface{})["login"].(string) + "のstar\n" + webhookInfos["repository"].(map[string]interface{})["html_url"].(string)
				for _, i := range sendGroups {
					bot.DingQQBot.SendGroupMessage(tool.StringToUint32(i), []message.IMessageElement{message.NewText(messageSend)})
				}
			}
		}
	}
	return JsonMessage(ctx, 0, "ok")
}
