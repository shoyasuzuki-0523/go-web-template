package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ConfigType struct {
	PORT      int
	DB_DSN    string
	FRONT_URL string
}

var Config ConfigType

func Init() {
	viper.SetConfigName(Env.Name())
	viper.SetConfigType("env")
	viper.AddConfigPath("common/config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		panic("Error loading .env file: ")
	}

	if err := viper.Unmarshal(&Config); err != nil {
		panic("Error Unmarshal .env file")
	}
}
