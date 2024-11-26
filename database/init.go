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
	DB.AutoMigrate(&KeyWordTable{}, &OpenTable{}, &StopTable{}, &MessageTable{})
	DB.Delete(&MessageTable{}, "1=1")
}
