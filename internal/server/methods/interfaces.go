package methods

import "github.com/gin-gonic/gin"

type Service struct {
	service *Method
}

type Method interface {
	GetStatistic(c *gin.Context)
	GetResourcesInfo(c *gin.Context)
	AddNewEmployee(c *gin.Context)
	AddNewOwner(c *gin.Context)
}
