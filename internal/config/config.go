package config

import (
	"errors"
	"os"
	"sync"

	"github.com/antikuz/xserver_exporter/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Url      string `yaml:"url" env:"URL" env-required:"true"`
	Login    string `yaml:"login" env:"LOGIN" env-required:"true"`
	Passwd   string `yaml:"passwd" env:"PASSWD" env-required:"true"`
	Insecure bool   `yaml:"insecure" env:"INSECURE" env-required:"true"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		var err error

		logger := logging.GetLogger()
		logger.Info("read exporter configuration")

		instance = &Config{}
		help, _ := cleanenv.GetDescription(instance, nil)

		if _, err = os.Stat("config.yaml"); errors.Is(err, os.ErrNotExist) {
			err = cleanenv.ReadEnv(instance)
		} else {
			err = cleanenv.ReadConfig("config.yaml", instance)
		}

		if err != nil {
			logger.Info(help)
			logger.Fatal(err)
		}

	})
	return instance
}
