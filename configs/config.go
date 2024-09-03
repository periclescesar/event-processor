package configs

import "github.com/spf13/viper"

func InitConfigs() {
	viper.AutomaticEnv()
	buildAppConfig()
	buildRabbitmqConfig()
	buildMongodbConfig()
}
