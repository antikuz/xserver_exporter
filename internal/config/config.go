package config

import (
    "os"

    "github.com/antikuz/xserver_exporter/pkg/logging"
    "github.com/spf13/pflag"
    "github.com/spf13/viper"
)

type Config struct {
    Url      string
    Login    string
    Passwd   string
    Insecure bool
    LogLevel string `mapstructure:"log-level"`
}

var instance *Config

func GetConfig() *Config {
    logger := logging.GetLogger()
    pflag.StringP("url", "u","", "Xserver configuration file path.")
    pflag.StringP("login", "l","", "User account to authenticate.")
    pflag.StringP("passwd", "p","", "User account password.")
    pflag.BoolP("insecure", "i", false, "Allow insecure server connections when using SSL")
    pflag.String("log-level", "info", "the maximum level of messages that should be logged. (possible values: debug, info, warn, error)")
    pflag.StringP("config-file", "c","", "xserver configuration file path.")
    pflag.BoolP("help", "h", false, "Show help.")
    pflag.CommandLine.SortFlags = false
    pflag.Parse()
    
    help, err := pflag.CommandLine.GetBool("help")
    if err != nil {
        logger.Fatalf("Failed to parse help flag, due to err: %v", err)
    }

    if help {
        pflag.Usage()
        os.Exit(0)
    }

    err = viper.BindPFlags(pflag.CommandLine)
    if err != nil {
        logger.Fatalf("Failed to viper.BindPFlags, due to err: %v", err)
    }
    err = viper.BindEnv("URL")
    if err != nil {
        logger.Fatalf("Failed to viper.BindEnv, due to err: %v", err)
    }
    err = viper.BindEnv("LOGIN")
    if err != nil {
        logger.Fatalf("Failed to viper.BindEnv, due to err: %v", err)
    }
    err = viper.BindEnv("PASSWD")
    if err != nil {
        logger.Fatalf("Failed to viper.BindEnv, due to err: %v", err)
    }
    err = viper.BindEnv("INSECURE")
    if err != nil {
        logger.Fatalf("Failed to viper.BindEnv, due to err: %v", err)
    }
    err = viper.BindEnv("LOGLEVEL")
    if err != nil {
        logger.Fatalf("Failed to viper.BindEnv, due to err: %v", err)
    }

    logger.Info("read exporter configuration")
    if viper.GetString("config-file") != "" {
        viper.SetConfigFile(viper.GetString("config-file"))
        if err := viper.ReadInConfig(); err != nil {
            logger.Fatalf("fatal error config file: %v", err)
        }
    }

    instance = &Config{}
    err = viper.Unmarshal(instance)
    if err != nil {
        logger.Fatalf("unable to decode config into struct, %v", err)
    }
    
    logger.Printf("%+v", instance)
    return instance
}