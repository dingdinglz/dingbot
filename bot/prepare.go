package bot

import (
	"os"
	"path/filepath"
	"time"

	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/client/auth"
)

func PrepareBot() {
	rootPath, _ := os.Getwd()
	DingQQBot = client.NewClientEmpty()
	LoadDeviceInfo()
	DingQQBot.UseDevice(DeviceInfo)
	DingQQBot.UseVersion(auth.AppList["linux"]["3.2.15-30366"])
	DingQQBot.AddSignServer("https://sign.lagrangecore.org/api/sign/30366")

	DingQQBot.GroupMessageEvent.Subscribe(BotGroupMessageEvent)
	DingQQBot.PrivateMessageEvent.Subscribe(BotPrivateMessageEvent)
	DingQQBot.GroupRecallEvent.Subscribe(GroupRecallEvent)

	qrcode, _, _ := DingQQBot.FetchQRCodeDefault()
	os.WriteFile(filepath.Join(rootPath, "data", "qrcode.png"), qrcode, os.ModePerm)
	for {
		retCode, _ := DingQQBot.GetQRCodeResult()
		if retCode.Waitable() {
			time.Sleep(1 * time.Second)
		}
		if retCode.Success() {
			break
		}
	}

	os.Remove(filepath.Join(rootPath, "data", "qrcode.png"))
	_, e := DingQQBot.QRCodeLogin()
	if e != nil {
		panic(e)
	}

}
