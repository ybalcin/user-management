package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	DatabaseSettings struct {
		DatabaseURI  string `yaml:"DatabaseURI" json:"DatabaseURI"`
		DatabaseName string `yaml:"DatabaseName" json:"DatabaseName"`
	} `yaml:"DatabaseSettings"`
}

func Read() (*Config, error) {
	dir, _ := os.Getwd()
	file := fmt.Sprintf("%s/%s", dir, "config/user-management.yaml")

	v := viper.GetViper()
	v.SetConfigType("yaml")
	v.SetConfigFile(file)

	if e := v.ReadInConfig(); e != nil {
		return nil, e
	}

	cfg := new(Config)

	if e := v.Unmarshal(cfg); e != nil {
		return nil, e
	}

	return cfg, nil
}
