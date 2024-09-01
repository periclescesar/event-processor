package configs

import "github.com/spf13/viper"

func InitConfigs() {
	viper.AutomaticEnv()
	buildRabbitmqConfig()
	buildMongodbConfig()
}
