package config

import (
	"sync"

	"github.com/antikuz/xserver_exporter/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Url          string `yaml:"url" env:"URL" env-required:"true"`
	Login        string `yaml:"login" env:"LOGIN" env-required:"true"`
	Passwd       string `yaml:"passwd" env:"PASSWD" env-required:"true"`
	InsecureSkip bool   `yaml:"insecureSkip" env:"INSECURESKIP" env-required:"true"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read exporter configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yaml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}