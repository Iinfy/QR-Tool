package ui

import (
	"fmt"
	"qrgen/app"

	hook "github.com/robotn/gohook"
	v "github.com/spf13/viper"
)

func StartKeyboardHook() {
	add()

}

func EndKeyboardHook() {
	hook.End()
	fmt.Println("--- Keyboard hook ended ---")
}

func add() {
	fmt.Println("--- Keyboard hook started ---")
	hook.Register(hook.KeyDown, v.GetStringSlice("gen.hotkey"), func(e hook.Event) {
		fmt.Println("qr generation")
		app.UrlToQR()
	})

	hook.Register(hook.KeyDown, v.GetStringSlice("scanner.hotkey"), func(e hook.Event) {
		fmt.Println("qr scanning")
		app.QRToUrl()
	})

	s := hook.Start()
	<-hook.Process(s)
}
