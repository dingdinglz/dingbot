package bot

import (
	"strconv"
	"strings"

	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/message"
	"github.com/dingdinglz/dingbot/appconfig"
	"github.com/dingdinglz/dingbot/database"
	"github.com/robertkrimen/otto"
)

func BotGroupMessageEvent(client *client.QQClient, event *message.GroupMessage) {
	if !database.OpenHave(event.GroupUin, "group") {
		return
	}

	messageText := event.ToString()
	if messageText == "dingbot" {
		client.SendGroupMessage(event.GroupUin, []message.IMessageElement{message.NewText("dingbot version:" + appconfig.Version + "\nwebsite:https://github.com/dingdinglz/dingbot")})
	}
	// 存表，以备查询
	database.MessageInsert("group", messageText, event.OriginalObject.ContentHead.Sequence.Unwrap())

	// 插件
	{
		RunPlugin(func(vm *otto.Otto) {
			vm.Set("a", strconv.Itoa(int(event.GroupUin)))
			vm.Set("b", messageText)
			vm.Set("c", strconv.Itoa(int(event.Sender.Uin)))
			vm.Run("GroupMessageEvent(a,b,c)")
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
			// 发言者是普通成员
			if groupMembers[client.Uin].Permission != 0 {
				// 机器人是管理
			}
		} else {
			// 发言者是管理
			var ReplyFlag = false
			var ReplySeq uint32
			var ReplySender uint32
			for _, i := range event.GetElements() {
				if i.Type() == message.Reply {
					ReplyFlag = true
					ReplySeq = i.(*message.ReplyElement).ReplySeq
					ReplySender = i.(*message.ReplyElement).SenderUin
				}
				if i.Type() == message.Text {
					if strings.ReplaceAll(i.(*message.TextElement).Content, " ", "") == "撤回" {
						if ReplyFlag {
							if ReplySender == client.Uin {
								client.RecallGroupMessage(event.GroupUin, ReplySeq)
							} else {
								if groupMembers[client.Uin].Permission > 0 && groupMembers[ReplySender].Permission == 0 {
									client.RecallGroupMessage(event.GroupUin, ReplySeq)
								}
							}
						}
					}
				}
			}
		}
	}
}
