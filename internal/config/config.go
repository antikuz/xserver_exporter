package config

import (
	"log"
	"sync"

	"github.com/antikuz/xserver_exporter/pkg/logging"
	"github.com/spf13/viper"
)

type Config struct {
	Url      string
	Login    string
	Passwd   string
	Insecure bool
	LogLevel string
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	logger := logging.GetLogger()
	logger.Info("read exporter configuration")

	viper.BindEnv("url")
	viper.BindEnv("login")
	viper.BindEnv("passwd")
	viper.BindEnv("insecure")
	viper.BindEnv("loglevel")

	viper.SetDefault("loglevel", "Info")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Warn("No Config file found, loaded config from Environment - Default path ./config.yaml")
		} else {
			logger.Fatalf("fatal error config file: %v", err)
		}
	}

	instance = &Config{}
	err := viper.Unmarshal(instance)
	if err != nil {
		log.Fatalf("unable to decode config into struct, %v", err)
	}

	return instance
}