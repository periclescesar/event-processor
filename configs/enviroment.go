package configs

import (
	"github.com/spf13/viper"
	"log"
)

func InitEnv(configFile string) {
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading configs file %s", err)
	}
}
