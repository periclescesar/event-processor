package configs

import "github.com/spf13/viper"

type mongodb struct {
	Uri string
}

var Mongodb *mongodb

func buildMongodbConfig() {
	var uri = viper.Get("MONGODB_CONNECTION_URI").(string)
	Mongodb = &mongodb{Uri: uri}
}
