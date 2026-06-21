package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ApiKey  string
	AppPort string
	DBHost  string
	DBUser  string
	DBPass  string
	DBName  string
	LogLevel string
}

func LoadConfig() *Config {
	
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return &Config{
		ApiKey:   v.GetString("api_key"),
		AppPort:  v.GetString("app.port"),
		DBHost:   v.GetString("database.host"),
		DBUser:   v.GetString("database.user"),
		DBPass:   v.GetString("database.pass"),
		DBName:   v.GetString("database.name"),
		LogLevel: v.GetString("log.level"),
	}
}
