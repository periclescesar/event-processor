package configs

import "github.com/spf13/viper"

type rabbitmq struct {
	Uri string
}

var Rabbitmq *rabbitmq

func buildRabbitmqConfig() {
	var uri = viper.Get("RABBITMQ_CONNECTION_URI").(string)
	Rabbitmq = &rabbitmq{Uri: uri}
}