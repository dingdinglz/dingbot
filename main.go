package main

import (
	"fmt"

	"github.com/dingdinglz/dingbot/route"
)

func main() {
	fmt.Println("dingbot " + route.Version + " Copyright © dinglz 2024")
	ConfigInit()
	ServerInit()
	ServerRun()
}
