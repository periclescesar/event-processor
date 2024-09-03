package configs

import "github.com/spf13/viper"

type app struct {
	LogLevel string
}

var App *app

func buildAppConfig() {
	level := viper.GetString("LOG_LEVEL")
	App = &app{LogLevel: level}
}
