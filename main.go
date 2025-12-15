package main

import (
	"qrgen/config"
	"qrgen/ui"
	"time"

	"github.com/gen2brain/beeep"
	"golang.design/x/clipboard"
)

func main() {
	config.ImportConfig()
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	beeep.AppName = "QR Tool"
	go ui.StartKeyboardHook()
	defer ui.EndKeyboardHook()
	time.Sleep(1 * time.Second)
	ui.LaunchTrayMenu()

}
