package server

import (
	"fmt"
	"github.com/spf13/viper"
)

type Service struct {
	service *WAFConfig
}

type WAFConfig interface {
	LoadConfig(path string) (config Config, err error)
}

type Config struct {
	POSTGRESQL_CONNSTRING string `mapstructure:"POSTGRESQL_CONNSTRING"`
}

//LoadConfig функция обработки конфига
func (service *Service) LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("err in config.go: ", err.Error())
	}

	err = viper.Unmarshal(&config)
	return
}
