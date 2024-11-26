package database

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DatabaseInitTest() {
	DB, _ = gorm.Open(sqlite.Open("../data/data.db"))
	DB.AutoMigrate(&KeyWordTable{}, &OpenTable{}, &StopTable{}, &MessageTable{}, &PluginData{})
	DB.Delete(&MessageTable{}, "1=1")
}

func TestPluginInsert(t *testing.T) {
	DatabaseInitTest()
	PluginDataSet("test", "test", "2")
}
