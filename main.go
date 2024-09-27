package main

import "fmt"

func main() {
	fmt.Println("dingbot " + Version + " Copyright © dinglz 2024")
	ConfigInit()
	ServerInit()
	ServerRun()
}
