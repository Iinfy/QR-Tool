package ui

import (
	"fmt"
	"qrgen/app"

	hook "github.com/robotn/gohook"
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
	hook.Register(hook.KeyDown, []string{"ctrl", "alt", "q"}, func(e hook.Event) {
		fmt.Println("ctrl-alt-q", "qr generation")
		app.UrlToQR()
	})

	hook.Register(hook.KeyDown, []string{"ctrl", "alt", "e"}, func(e hook.Event) {
		fmt.Println("ctrl-alt-q", "qr scanning")
		app.QRToUrl()
	})

	s := hook.Start()
	<-hook.Process(s)
}
