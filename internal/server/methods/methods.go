package methods

import "github.com/gin-gonic/gin"

type Service struct {
	service *Method
}

type Method interface {
	GetResourcesInfo(c *gin.Context)
}
