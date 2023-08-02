package server

import (
	"WAF_Analytics/internal/server/resource"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Service struct {
	service *Create
}

type Create interface {
	CreateServer()
}

func (service *Service) CreateServer() {
	method := resource.PgService{}
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers, Authorization"},
	}))
	router.POST("/api/login", method.Login)
	router.GET("/api/stats", method.GetStat)
	router.GET("/api/get-res-stat", method.GetResStat)
	router.POST("/api/add-employee", method.AddEmployee)
	router.POST("/api/add-owner", method.AddOwner)
	router.POST("/api/add-resource", method.AddResource)
	router.POST("/api/get-resource-by-owner", method.FindResourceByOwner)
	router.GET("/api/get-info-about-owners", method.GetInformationAboutOwner)
	router.POST("/api/delete-owner", method.DeleteOwner)
	router.POST("/api/delete-resource", method.DeleteResource)
	router.POST("/api/update-resource", method.UpdateResource)
	router.GET("/api/get-week-stat", method.GetWeekStat)
	router.POST("/api/check-resource", method.CheckResource)
	//router.GET("/api/get-common-info", method.CounterCommonInfo)
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("err: ", err)
	}
}
