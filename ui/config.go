package ui

import (
	"log"

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
