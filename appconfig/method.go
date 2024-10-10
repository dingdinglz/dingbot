package appconfig

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func SaveBotConfig(username string, password string) {
	BotConfigVar.Username = username
	BotConfigVar.Password = password
	res, _ := json.Marshal(BotConfigVar)
	rootPath, _ := os.Getwd()
	os.WriteFile(filepath.Join(rootPath, "data", "bot.json"), res, os.ModePerm)
}
