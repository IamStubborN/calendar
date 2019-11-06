package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Logger  Logger  `mapstructure:"logger"`
		Storage Storage `mapstructure:"storage"`
	}

	Logger struct {
		Level string `mapstructure:"level"`
	}

	Storage struct {
		Provider string `mapstructure:"provider"`
		DSN      string `mapstructure:"dsn"`
		Retry    int    `mapstructure:"retry"`
	}
)

func LoadConfig() (*Config, error) {
	var config *Config
	viper.SetDefault("CALENDAR_CONFIG", "config.yaml")
	viper.SetConfigFile(viper.GetString("CALENDAR_CONFIG"))
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "unable to read config with filepath")
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal config to struct")
	}

	return config, nil
}
