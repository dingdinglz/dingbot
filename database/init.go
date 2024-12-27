package database

import (
	"github.com/dingdinglz/dingbot/appconfig"
	"github.com/dingdinglz/dingbot/tool"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DatabaseInit() {
	var err error
	switch appconfig.AppConfigVar.DatabaseType {
	case "sqlite":
		DB, err = gorm.Open(sqlite.Open(appconfig.AppConfigVar.Source))
	case "mysql":
		DB, err = gorm.Open(mysql.Open(appconfig.AppConfigVar.Source))
	default:
		panic("database type error")
	}
	if err != nil {
		tool.ErrorOut("database", "open error", err)
		panic(err.Error())
	}
	DB.AutoMigrate(&KeyWordTable{}, &OpenTable{}, &StopTable{}, &MessageTable{}, &PluginData{}, &StarTable{}, &GithubHookTable{}, &UserTable{})

	// 清空message表的内容，message表用于存seq与message对应，便于取撤回内容
	DB.Delete(&MessageTable{}, "1=1")
}
