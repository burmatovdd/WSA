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
}
