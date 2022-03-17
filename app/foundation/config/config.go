package config

import (
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	PGHostname string `mapstructure:"pghostname"`
	PGPort     string `mapstructure:"pgport"`
	PGUser     string `mapstructure:"pguser"`
	PGPassword string `mapstructure:"pgpassword"`
	PGDatabase string `mapstructure:"pgdatabase"`
	Port       string `mapstructure:"port"`
	SigningKey string `mapstructure:"signingkey"`
}

var Conf *Config

func getConf(logger *zap.Logger) *Config {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		logger.Fatal("Failed to read config file", zap.Error(err))
		os.Exit(3)
	}

	conf := &Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		logger.Fatal("unable to decode into config struct, %v", zap.Error(err))
		os.Exit(4)
	}

	return conf
}

// Initialization of config
func InitConfig(logger *zap.Logger) {
	Conf = getConf(logger)
	SigningKey = []byte(Conf.SigningKey)
}
