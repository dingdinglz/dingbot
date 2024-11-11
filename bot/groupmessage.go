package bot

import (
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/message"
	"github.com/dingdinglz/dingbot/appconfig"
	"github.com/dingdinglz/dingbot/database"
	"github.com/dingdinglz/dingbot/tool"
	"github.com/robertkrimen/otto"
)

func AddVMFuncs(Vm *otto.Otto) {
	Vm.Set("SendGroupMessage", func(call otto.FunctionCall) otto.Value {
		DingQQBot.SendGroupMessage(tool.StringToUint32(call.Argument(0).String()), []message.IMessageElement{message.NewText(call.Argument(1).String())})
		return otto.Value{}
	})
	Vm.Set("SendPrivateMessage", func(call otto.FunctionCall) otto.Value {
		DingQQBot.SendPrivateMessage(tool.StringToUint32(call.Argument(0).String()), []message.IMessageElement{message.NewText(call.Argument(1).String())})
		return otto.Value{}
	})
}

func BotGroupMessageEvent(client *client.QQClient, event *message.GroupMessage) {
	if !database.OpenHave(event.GroupUin, "group") {
		return
	}

	messageText := event.ToString()
	if messageText == "dingbot" {
		client.SendGroupMessage(event.GroupUin, []message.IMessageElement{message.NewText("dingbot version:" + appconfig.Version + "\nwebsite:https://github.com/dingdinglz/dingbot")})
	}

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
				vm.Set("a", strconv.Itoa(int(event.GroupUin)))
				vm.Set("b", messageText)
				vm.Set("c", strconv.Itoa(int(event.Sender.Uin)))
				vm.Run("GroupMessageEvent(a,b,c)")
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
					client.SendGroupMessage(event.GroupUin, []message.IMessageElement{message.NewGroupReply(event), message.NewText(i.Text)})
				}
			} else {
				if strings.Contains(messageText, i.Key) {
					client.SendGroupMessage(event.GroupUin, []message.IMessageElement{message.NewGroupReply(event), message.NewText(i.Text)})
				}
			}
		}
	}

	{
		groupMembers, _ := DingQQBot.GetGroupMembersData(event.GroupUin)
		if groupMembers[event.Sender.Uin].Permission == 0 {
			if groupMembers[client.Uin].Permission != 0 {
			}
		}
	}
}
