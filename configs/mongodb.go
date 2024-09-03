package configs

import "github.com/spf13/viper"

type mongodb struct {
	URI    string
	DBname string
}

var Mongodb *mongodb

func buildMongodbConfig() {
	var uri = viper.Get("MONGODB_CONNECTION_URI").(string)
	var dbName = viper.Get("MONGODB_DB_NAME").(string)
	Mongodb = &mongodb{URI: uri, DBname: dbName}
}
