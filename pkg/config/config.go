package config

import (
	"fmt"
	"github.com/google/wire"
	"github.com/spf13/viper"
)

type Path string

func New(configPath Path) (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigFile(string(configPath))
	fmt.Printf("Config path: %s\n", configPath)
	v.WatchConfig()
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config path: %s \n", err))
	}
	return v, err
}

var ProvideSet = wire.NewSet(New)
