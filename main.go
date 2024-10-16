package main

import (
	"fmt"

	"github.com/dingdinglz/dingbot/appconfig"
	"github.com/dingdinglz/dingbot/database"
)

func main() {
	fmt.Println("dingbot " + appconfig.Version + " Copyright © dinglz 2024")
	appconfig.ConfigInit()
	database.DatabaseInit()
	ServerInit()
	ServerRun()
}
