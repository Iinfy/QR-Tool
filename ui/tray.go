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
	systray.SetTitle("QR Tool")
	systray.SetTooltip("QR Tool")

	generationMenu := systray.AddMenuItem("QR Generation", "QR generation settings")
	generationEnabled := generationMenu.AddSubMenuItemCheckbox("Enabled", "Enabled", true)

	scanningMenu := systray.AddMenuItem("QR Scanner", "QR scanner setiings")
	scanEnabled := scanningMenu.AddSubMenuItemCheckbox("Enabled", "Enabled", true)

	systray.AddSeparator()

	mQuit := systray.AddMenuItem("Quit", "Exit the application")

	go func() {
		for {
			select {
			case <-mQuit.ClickedCh:
				systray.Quit()
				return

			case <-generationEnabled.ClickedCh:
				if generationEnabled.Checked() {
					generationEnabled.Uncheck()
				} else {
					generationEnabled.Check()
				}

			case <-scanEnabled.ClickedCh:
				if scanEnabled.Checked() {
					scanEnabled.Uncheck()
				} else {
					scanEnabled.Check()
				}
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
