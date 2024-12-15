package route

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GithubWebhookRoute(ctx *fiber.Ctx) error {
	eventName := ctx.Get("X-GitHub-Event")
	var webhookInfos map[string]interface{}
	json.Unmarshal(ctx.Body(), &webhookInfos)
	repositoryName := webhookInfos["repository"].(map[string]interface{})["full_name"]
	fmt.Println(repositoryName)
	switch eventName {
	case "push":
		//bot.DingQQBot.SendGroupMessage(11, []message.IMessageElement{message.NewText("11")})
	case "star":
	}
	return JsonMessage(ctx, 0, "ok")
}
