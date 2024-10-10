package appconfig

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/dingdinglz/dingbot/tool"
)

func ConfigInit() {
	rootPath, _ := os.Getwd()
	tool.DictoryCreateN(filepath.Join(rootPath, "data"))
	if !tool.FileIsExsits(filepath.Join(rootPath, "data", "config.json")) {
		Init = true
		AppConfigVar.Port = 7000
	} else {
		Init = false
		jsonText, _ := os.ReadFile(filepath.Join(rootPath, "data", "config.json"))
		json.Unmarshal(jsonText, &AppConfigVar)
	}

	if !tool.FileIsExsits(filepath.Join(rootPath, "data", "bot.json")) {
		BotConfigVar.Username = ""
		BotConfigVar.Password = ""
	} else {
		jsonText, _ := os.ReadFile(filepath.Join(rootPath, "data", "bot.json"))
		json.Unmarshal(jsonText, &BotConfigVar)
	}
}
