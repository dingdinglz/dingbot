package main

import (
	"fmt"

	"github.com/dingdinglz/dingbot/appconfig"
)

func main() {
	fmt.Println("dingbot " + appconfig.Version + " Copyright © dinglz 2024")
	appconfig.ConfigInit()
	ServerInit()
	ServerRun()
}
