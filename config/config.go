package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	API_PORT string
}

var Cfg *Config

func InitConfig() {
	cfg := &Config{}

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}

	_ = viper.Unmarshal(cfg)
	Cfg = cfg
	fmt.Println(Cfg)
}
