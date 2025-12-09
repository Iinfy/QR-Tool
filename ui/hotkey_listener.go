package ui

import (
	"fmt"
	"qrgen/app"

	hook "github.com/robotn/gohook"
)

func StartKeyboardHook() {
	add()

	low()
}

func EndKeyboardHook() {
	hook.End()
}

func add() {
	fmt.Println("--- Keyboard hook started ---")
	hook.Register(hook.KeyDown, []string{"ctrl", "alt", "q"}, func(e hook.Event) {
		fmt.Println("ctrl-alt-q", "qr generation")
		app.UrlToQR()
	})

	hook.Register(hook.KeyDown, []string{"ctrl", "alt", "e"}, func(e hook.Event) {
		fmt.Println("ctrl-alt-q", "qr scanning")
	})

	s := hook.Start()
	<-hook.Process(s)
}

func low() {
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		fmt.Println("hook: ", ev)
	}
}
