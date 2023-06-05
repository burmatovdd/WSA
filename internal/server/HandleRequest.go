package server

import (
	"WAF_Analytics/internal/server/methods"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Service struct {
	service *HandleRequest
}

type HandleRequest interface {
	HandleRequest()
}

func (service *Service) HandleRequest() {
	method := methods.Service{}
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers, Authorization"},
	}))
	router.GET("/api/getStatistic", method.GetStatistic)
	router.GET("/api/getResourcesInfo", method.GetResourcesInfo)
	router.GET("/api/addNewEmployee", method.AddNewEmployee)
	router.GET("/api/addNewOwner", method.AddNewOwner)
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("err: ", err)
	}
}
