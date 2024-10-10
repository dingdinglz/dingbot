package bot

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/client/auth"

	"github.com/dingdinglz/dingbot/appconfig"
	"github.com/dingdinglz/dingbot/tool"
)

func PrepareBot() {
	rootPath, _ := os.Getwd()
	i, _ := strconv.ParseUint(appconfig.BotConfigVar.Username, 10, 32)
	DingQQBot = client.NewClient(uint32(i), auth.AppList["linux"]["3.1.2-13107"], "https://sign.lagrangecore.org/api/sign")
	LoadDeviceInfo()
	DingQQBot.UseDevice(DeviceInfo)
	if tool.FileIsExsits(filepath.Join(rootPath, "data", "sig.bin")) {
		FirstLogin = false
		res, _ := os.ReadFile(filepath.Join(rootPath, "data", "sig.bin"))
		sig, err := auth.UnmarshalSigInfo(res, true)
		if err != nil {
			FirstLogin = true
		} else {
			DingQQBot.UseSig(sig)
		}
	} else {
		FirstLogin = true
	}
	DingQQBot.GroupMessageEvent.Subscribe(BotGroupMessageEvent)
	err := DingQQBot.Login("", filepath.Join(rootPath, "data", "qrcode.png"))
	if err != nil {
		fmt.Println(err.Error())
	}
	sigBin, e := DingQQBot.Sig().Marshal()
	os.WriteFile(filepath.Join(rootPath, "data", "sig.bin"), sigBin, os.ModePerm)
	if e != nil {
		fmt.Println(e.Error())
	}
	if tool.FileIsExsits(filepath.Join(rootPath, "data", "qrcode.png")) {
		os.Remove(filepath.Join(rootPath, "data", "qrcode.png"))
	}
}
