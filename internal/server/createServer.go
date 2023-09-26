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

	api := router.Group("/api", method.UserIdentity)
	{
		api.GET("/chart-statistic", method.GetShortStat)
		api.GET("/week-statistic", method.GetWeekStat)
		api.GET("/general-statistic", method.GetGeneralStat)
		api.GET("/certificates", method.GetCertificates)
		api.GET("/statistic", method.GetStatistic)
		api.GET("/user-info", method.GetUserInfo)

		api.POST("/add-resource", method.AddResource)
		api.POST("/check-resource", method.CheckResource)
		api.POST("/delete-resource", method.DeleteResource)
		api.POST("/update-resource", method.UpdateResource)
	}
	router.POST("/login", method.Login)

	err := router.Run(":8080")
	if err != nil {
		fmt.Println("err: ", err)
	}
}
