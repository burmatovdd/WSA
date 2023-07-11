package resource

import (
	"github.com/gin-gonic/gin"
)

type PgService struct {
	service *Service
}

type Service interface {
	GetStat(c *gin.Context)
	GetResStat(c *gin.Context)
	AddResource(c *gin.Context)
	AddEmployee(c *gin.Context)
	AddOwner(c *gin.Context)
	FindResourceByOwner(c *gin.Context)
	GetInformationAboutOwner(c *gin.Context)
	DeleteOwner(c *gin.Context)
	DeleteResource(c *gin.Context)
	UpdateResource(c *gin.Context)
	GetWeekStat(c *gin.Context)
	Login(c *gin.Context)
}
