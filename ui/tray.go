package ui

import (
	"fmt"
	"os"
	"qrgen/app"
	"qrgen/config"
	"qrgen/utils"

	"github.com/getlantern/systray"
	v "github.com/spf13/viper"
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
	generationEnabled := generationMenu.AddSubMenuItemCheckbox("Enabled", "Enabled", v.GetBool("gen.enabled"))

	scanningMenu := systray.AddMenuItem("QR Scanner", "QR scanner setiings")
	scanEnabled := scanningMenu.AddSubMenuItemCheckbox("Enabled", "Enabled", v.GetBool("scanner.enabled"))

	mainMonitor := systray.AddMenuItem("Main monitor", "Select main monitor")
	monitor1 := mainMonitor.AddSubMenuItemCheckbox("Monitor 1", "Set 1 monitor as main", utils.IsMainDisplay(1))
	monitor2 := mainMonitor.AddSubMenuItemCheckbox("Monitor 2", "Set 2 monitor as main", utils.IsMainDisplay(2))
	monitor3 := mainMonitor.AddSubMenuItemCheckbox("Monitor 3", "Set 3 monitor as main", utils.IsMainDisplay(3))

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
					v.Set("gen.enabled", false)
					app.SetGeneratorEnabled(false)
					config.SaveConfig()
				} else {
					generationEnabled.Check()
					app.SetGeneratorEnabled(true)
					v.Set("gen.enabled", true)
					config.SaveConfig()
				}

			case <-scanEnabled.ClickedCh:
				if scanEnabled.Checked() {
					scanEnabled.Uncheck()
					app.SetScannerEnabled(false)
					v.Set("scanner.enabled", false)
					config.SaveConfig()
				} else {
					scanEnabled.Check()
					app.SetScannerEnabled(true)
					v.Set("scanner.enabled", true)
					config.SaveConfig()
				}

			case <-monitor1.ClickedCh:
				monitor1.Check()
				monitor2.Uncheck()
				monitor3.Uncheck()
				utils.SetActiveDisplay(1)
				v.Set("scanner.monitor", 1)
				config.SaveConfig()

			case <-monitor2.ClickedCh:
				monitor1.Uncheck()
				monitor2.Check()
				monitor3.Uncheck()
				utils.SetActiveDisplay(2)
				v.Set("scanner.monitor", 2)
				config.SaveConfig()

			case <-monitor3.ClickedCh:
				monitor1.Uncheck()
				monitor2.Uncheck()
				monitor3.Check()
				utils.SetActiveDisplay(3)
				v.Set("scanner.monitor", 3)
				config.SaveConfig()
			}

		}
	}()
}

func onExit() {
	fmt.Println("Application exited gracefully.")
}
