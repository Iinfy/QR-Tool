package ui

import (
	"fmt"
	"os"
	"time"

	"github.com/getlantern/systray"
)

func getIcon() []byte {
	icoBytes, err := os.ReadFile("qrlogo.ico")
	if err != nil {
		panic(err)
	}
	return icoBytes
}

func LaunchTrayMenu() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(getIcon())
	systray.SetTitle("QRT")
	systray.SetTooltip("QR Tool")

	mOpen := systray.AddMenuItem("Open Dashboard", "Open the application UI")

	systray.AddSeparator()

	mToggle := systray.AddMenuItemCheckbox("Service Running", "Toggle the main service", true)

	mQuit := systray.AddMenuItem("Quit", "Exit the application")

	go func() {
		for {
			select {
			case <-mOpen.ClickedCh:
				fmt.Println("Opening application dashboard...")

			case <-mToggle.ClickedCh:
				if mToggle.Checked() {
					mToggle.Uncheck()
					fmt.Println("Service Stopped")
				} else {
					mToggle.Check()
					fmt.Println("Service Started")
				}

			case <-mQuit.ClickedCh:
				fmt.Println("Quit signal received, exiting...")
				systray.Quit()
				return
			}
		}
	}()

	go func() {
		for {
			systray.SetTitle(time.Now().Format("15:04:05"))
			time.Sleep(1 * time.Second)
		}
	}()
}

func onExit() {
	fmt.Println("Application exited gracefully.")
}
