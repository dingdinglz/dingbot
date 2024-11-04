package bot

import (
	"io/fs"
	"os"
	"path/filepath"
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
		rootPath, _ := os.Getwd()
		filepath.Walk(filepath.Join(rootPath, "data", "plugin"), func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if filepath.Ext(path) == ".js" {
				vm := otto.New()
				AddVMFuncs(vm)
				codeText, _ := os.ReadFile(path)
				vm.Run(string(codeText))
				vm.Set("a", strconv.Itoa(int(event.Sender.Uin)))
				vm.Set("b", messageText)
				vm.Run("PrivateMessageEvent(a,b)")
			}
			return nil
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
