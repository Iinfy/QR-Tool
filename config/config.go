package config

import (
	"log"
	"qrgen/app"
	"qrgen/utils"

	v "github.com/spf13/viper"
)

func ImportConfig() {
	v.SetConfigName("config")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		log.Println(err)
		setDefaultConfig()
	}
	utils.SetActiveDisplay(v.GetInt("scanner.monitor"))
	app.SetGeneratorEnabled(v.GetBool("gen.enabled"))
	app.SetScannerEnabled(v.GetBool("scanner.enabled"))

}

func setDefaultConfig() {
	v.Set("gen.enabled", true)
	genHotkey := [3]string{"ctrl", "alt", "q"}
	v.Set("gen.hotkey", genHotkey)

	v.Set("scanner.enabled", true)
	scanHotkey := [3]string{"ctrl", "alt", "e"}
	v.Set("scanner.hotkey", scanHotkey)
	v.Set("scanner.monitor", 1)

	v.WriteConfigAs("config.json")
}

func SaveConfig() {
	v.WriteConfigAs("config.json")
}
