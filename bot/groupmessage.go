package bot

import (
	"strings"

	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/message"
	"github.com/dingdinglz/dingbot/appconfig"
	"github.com/dingdinglz/dingbot/database"
)

func BotGroupMessageEvent(client *client.QQClient, event *message.GroupMessage) {
	if !database.OpenHave(event.GroupUin, "group") {
		return
	}

	messageText := event.ToString()
	if messageText == "dingbot" {
		client.SendGroupMessage(event.GroupUin, []message.IMessageElement{message.NewText("dingbot version:" + appconfig.Version + "\nwebsite:https://github.com/dingdinglz/dingbot")})
	}

	{
		// 关键词
		keywords := database.KeyWordGetAll()
		for _, i := range keywords {
			if i.Precise == "true" {
				if messageText == i.Key {
					client.SendGroupMessage(event.GroupUin, []message.IMessageElement{message.NewGroupReply(event), message.NewText(i.Text)})
				}
			} else {
				if strings.Contains(messageText, i.Key) {
					client.SendGroupMessage(event.GroupUin, []message.IMessageElement{message.NewGroupReply(event), message.NewText(i.Text)})
				}

			}
		}
	}
}
