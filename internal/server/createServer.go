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
		AllowHeaders: []string{"Content-Type, access-control-allow-origin, access-control-allow-headers, Authorization"},
	}))

	//api := router.Group("/api", method.UserIdentity)

	router.GET("/api/chart-statistic", method.GetStat)
	router.GET("/api/week-statistic", method.GetWeekStat)
	router.GET("/api/general-statistic", method.GetGeneralStat)
	router.GET("/api/certificates", method.GetCertificates)
	router.GET("/api/statistic", method.GetStatistic)

	router.POST("/api/add-resource", method.AddResource)
	router.POST("/api/check-resource", method.CheckResource)
	router.POST("/api/delete-resource", method.DeleteResource)
	router.POST("/api/update-resource", method.UpdateResource)

	router.POST("/login", method.Login)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("err: ", err)
	}
}
