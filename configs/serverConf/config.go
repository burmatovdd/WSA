package serverConf

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	POSTGRESQL_CONNSTRING string `mapstructure:"POSTGRESQL_CONNSTRING"`
}

var DefaultConfig Config

//LoadConfig функция обработки конфига
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("err in config.go: ", err.Error())
	}

	err = viper.Unmarshal(&config)
	DefaultConfig = config
	return
}
