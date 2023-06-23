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
	router.GET("/api/stats", method.GetStat)
	router.GET("/api/get-res-stat", method.GetResStat)
	router.GET("/api/add-employee", method.AddEmployee)
	router.GET("/api/add-owner", method.AddOwner)
	router.GET("/api/add-resource", method.AddResource)
	router.GET("/api/get-resource-by-owner", method.FindResourceByOwner)
	router.GET("/api/get-info-about-owners", method.GetInformationAboutOwner)
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("err: ", err)
	}
}
