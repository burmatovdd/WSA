package methods

import (
	"WAF_Analytics/configs/server"
	"WAF_Analytics/internal/server/postgresql"
	"github.com/gin-gonic/gin"
)

//GetResourcesInfo функция получения информации о всех ресурсах
func (service *Service) GetResourcesInfo(c *gin.Context) {
	method := postgresql.Service{}
	config := server.Service{}

	method.GetAllStat(config.LoadConfig("configs/server"))
}
