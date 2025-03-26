package bot

import (
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
		DeviceInfo, _ = auth.LoadOrSaveDevice(filepath.Join(rootPath, "data", "device.json"))
	}
}
