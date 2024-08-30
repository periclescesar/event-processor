package configs

func InitConfigs(configFile string) {
	InitEnv(configFile)
	buildRabbitmqConfig()
}
