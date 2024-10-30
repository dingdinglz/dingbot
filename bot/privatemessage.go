package bot

import (
	"strings"

	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/message"
	"github.com/dingdinglz/dingbot/database"
)

func BotPrivateMessageEvent(client *client.QQClient, event *message.PrivateMessage) {
	if !database.OpenHave(event.Sender.Uin, "private") {
		return
	}

	messageText := event.ToString()

	{
		// 关键词
		keywords := database.KeyWordGetAll()
		for _, i := range keywords {
			if i.Precise == "true" {
				if messageText == i.Key {
					client.SendPrivateMessage(event.Sender.Uin, []message.IMessageElement{message.NewText(i.Text)})
				}
			} else {
				if strings.Contains(messageText, i.Key) {
					client.SendPrivateMessage(event.Sender.Uin, []message.IMessageElement{message.NewText(i.Text)})
				}

			}
		}
	}
}
