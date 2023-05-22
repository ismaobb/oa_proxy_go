package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type env struct {
	Mode string  `mapstructure:"mode"`
	Port int     `mapstructure:"port"`
	KL   sqlConf `mapstructure:"kl"`
	KC   sqlConf `mapstructure:"kc"`
	KW   sqlConf `mapstructure:"kw"`
}

type sqlConf struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

func newEnv() *env {
	e := new(env)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	err := viper.Unmarshal(&e)
	if err != nil {
		log.Fatal(err)
	}

	return e
}
