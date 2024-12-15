package database

import (
	"os"
	"path/filepath"

	"github.com/dingdinglz/dingbot/tool"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DatabaseInit() {
	rootPath, _ := os.Getwd()
	var err error
	DB, err = gorm.Open(sqlite.Open(filepath.Join(rootPath, "data", "data.db")))
	if err != nil {
		tool.ErrorOut("database", "open error", err)
	}
	DB.AutoMigrate(&KeyWordTable{}, &OpenTable{}, &StopTable{}, &MessageTable{}, &PluginData{}, &StarTable{}, &GithubHookTable{})

	// 清空message表的内容，message表用于存seq与message对应，便于取撤回内容
	DB.Delete(&MessageTable{}, "1=1")
}
