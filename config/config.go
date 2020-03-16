package config

import (
	"log"
	"os"

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
		os.Exit(1)
	}

	if err := viper.Unmarshal(c); err != nil {
		log.Fatalln(err.Error())
		os.Exit(1)
	}
}
