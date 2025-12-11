package main

import (
	"qrgen/ui"

	"github.com/gen2brain/beeep"
	"golang.design/x/clipboard"
)

func main() {
	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	beeep.AppName = "QR Tool"
	go ui.StartKeyboardHook()
	defer ui.EndKeyboardHook()
	ui.LaunchTrayMenu()

}
