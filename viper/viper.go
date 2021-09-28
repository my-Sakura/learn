package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Host     string
	Port     int
	UserName string
	Password string
	Database string
	Charset  string
}

func New() *Config {
	return &Config{}
}

func (c *Config) ReadConfig() error {
	viper.SetConfigFile("./config.yml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return fmt.Errorf("config file not found %v\n", err)
		} else {
			return fmt.Errorf("config file could found %v\n", err)
		}
	}

	if err := viper.Unmarshal(&c); err != nil {
		return fmt.Errorf("unmarshal error: %v\n", err)
	}

	return nil
}
