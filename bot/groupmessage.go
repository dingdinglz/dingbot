package bot

import (
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/message"
	"github.com/dingdinglz/dingbot/appconfig"
)

func BotGroupMessageEvent(client *client.QQClient, event *message.GroupMessage) {
	messageText := event.ToString()
	if messageText == "dingbot" {
		client.SendGroupMessage(event.GroupUin, []message.IMessageElement{message.NewText("dingbot version:" + appconfig.Version + "\nwebsite:https://github.com/dingdinglz/dingbot")})
	}
}
