package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Port string
	}
	Mysql struct {
		Dsn         string
		MaxIdleConn int
		MaxOpenConn int
	}
	Redis struct {
		Addr     string
		Password string
		DB       int
	}
	Grpc struct {
		ServerPort string
		Protocol   string
		ClientAddr string
	}
}

func NewConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("setConfig  error: %v", err)
	}

	var AppConfig = &Config{}
	err = viper.Unmarshal(AppConfig)
	if err != nil {
		log.Fatalf("unmarshal config error: %v", err)
	}

	return AppConfig
}
