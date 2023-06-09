package main

import (
	"WAF_Analytics/configs/serverConf"
	"WAF_Analytics/internal/server"
	"log"
)

func main() {
	//загружаем конфиг на старте приложения
	_, err := serverConf.LoadConfig("configs/serverConf")
	if err != nil {
		log.Fatalln("error in config: ", err)
	}
	service := server.Service{}
	service.CreateServer()
}
