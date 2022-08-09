package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/antikuz/xserver_exporter/pkg/logging"
	"github.com/spf13/pflag"
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

	viper.BindEnv("URL")
	viper.BindEnv("LOGIN")
	viper.BindEnv("PASSWD")
	viper.BindEnv("INSECURE")
	viper.BindEnv("LOGLEVEL")

	viper.SetDefault("loglevel", "Info")

	if viper.GetString("config-file") != "" {
		viper.SetConfigFile(viper.GetString("config-file"))
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
	}

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
	log.Printf("%+v", instance)
	return instance
}

