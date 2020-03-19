package config

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	App struct {
		AppEnv string
	}
	Web struct {
		Host string
		Port string
	}
	Database struct {
		Driver     string
		SourcePath string
		SourceName string
		Username   string
		Password   string
	}
	Table struct {
		Music   string
		Artist  string
		Company string
	}
	Message struct {
		RouteNotFound    string
		MethodNotAllowed string
		RecordNotFound   string
		SaveError        string
	}
}

var Conf config

func ReadConf() {
	c := &Conf

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err.Error())
	}

	if err := viper.Unmarshal(c); err != nil {
		log.Fatalln(err.Error())
	}
}
