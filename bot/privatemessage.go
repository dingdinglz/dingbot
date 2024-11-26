package bot

import (
	"strconv"
	"strings"

	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/message"
	"github.com/dingdinglz/dingbot/database"
	"github.com/robertkrimen/otto"
)

func BotPrivateMessageEvent(client *client.QQClient, event *message.PrivateMessage) {
	if !database.OpenHave(event.Sender.Uin, "private") {
		return
	}

	messageText := event.ToString()

	// 插件
	{
		RunPlugin(func(vm *otto.Otto) {
			vm.Set("a", strconv.Itoa(int(event.Sender.Uin)))
			vm.Set("b", messageText)
			vm.Run("PrivateMessageEvent(a,b)")
		})
	}

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
