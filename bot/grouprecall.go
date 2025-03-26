package bot

import (
	"strconv"

	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/client/event"
	"github.com/dingdinglz/dingbot/database"
	"github.com/robertkrimen/otto"
)

func GroupRecallEvent(client *client.QQClient, event *event.GroupRecall) {
	messageText := database.MessageGetText("group", uint32(event.Sequence))
	RunPlugin(func(vm *otto.Otto) {
		vm.Set("a", messageText)
		vm.Set("b", strconv.Itoa(int(event.GroupUin)))
		vm.Set("c", strconv.Itoa(int(event.UserUin)))
		vm.Set("d", strconv.Itoa(int(event.OperatorUin)))
		vm.Run("GroupRecallEvent(a,b,c,d)")
	})
}
