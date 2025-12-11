package ui

import (
	"fmt"
	"os"
	"qrgen/utils"
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

	mainMonitor := systray.AddMenuItem("Main monitor", "Select main monitor")
	monitor1 := mainMonitor.AddSubMenuItemCheckbox("Monitor 1", "Set 1 monitor as main", true)
	monitor2 := mainMonitor.AddSubMenuItemCheckbox("Monitor 2", "Set 2 monitor as main", false)
	monitor3 := mainMonitor.AddSubMenuItemCheckbox("Monitor 3", "Set 3 monitor as main", false)

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

			case <-monitor1.ClickedCh:
				monitor1.Check()
				monitor2.Uncheck()
				monitor3.Uncheck()
				utils.SetActiveDisplay(1)

			case <-monitor2.ClickedCh:
				monitor1.Uncheck()
				monitor2.Check()
				monitor3.Uncheck()
				utils.SetActiveDisplay(2)

			case <-monitor3.ClickedCh:
				monitor1.Uncheck()
				monitor2.Uncheck()
				monitor3.Check()
				utils.SetActiveDisplay(3)
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
