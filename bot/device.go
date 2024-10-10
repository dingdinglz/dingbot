package bot

import (
	"encoding/json"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"

	"github.com/LagrangeDev/LagrangeGo/client/auth"
	"github.com/dingdinglz/dingbot/appconfig"
	"github.com/dingdinglz/dingbot/tool"
)

func generateRandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func LoadDeviceInfo() {
	rootPath, _ := os.Getwd()
	uin, _ := strconv.Atoi(appconfig.BotConfigVar.Username)
	if !tool.FileIsExsits(filepath.Join(rootPath, "data", "device.json")) {
		DeviceInfo = auth.NewDeviceInfo(uin)
		DeviceInfo.DeviceName = "Dingbot-" + generateRandomString(8)
		DeviceInfo.Save(filepath.Join(rootPath, "data", "device.json"))
	} else {
		var i DeviceStruct
		res, _ := os.ReadFile(filepath.Join(rootPath, "data", "device.json"))
		json.Unmarshal(res, &i)
		DeviceInfo = &auth.DeviceInfo{Guid: i.GUID, KernelVersion: i.KernelVersion, SystemKernel: i.SystemKernel, DeviceName: i.DeviceName}
	}
}
