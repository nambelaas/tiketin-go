package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigFile("./config/config.json")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
	}
}
