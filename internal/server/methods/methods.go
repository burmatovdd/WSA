package methods

import (
	"WAF_Analytics/configs/server"
	"WAF_Analytics/internal/server/postgresql"
	"fmt"
	"github.com/gin-gonic/gin"
)

//GetStatistic функция получения информации о всех ресурсах
func (service *Service) GetStatistic(c *gin.Context) {
	method := postgresql.Service{}
	config := server.Service{}

	method.GetAllStat(config.LoadConfig("configs/server"))
}

//GetResourcesInfo функция получения информации о ресурсе
func (service *Service) GetResourcesInfo(c *gin.Context) {
	method := postgresql.Service{}
	config := server.Service{}

	method.GetResourcesStat(config.LoadConfig("configs/server"))
}

//AddNewEmployee функция добавления нового пользователя
func (service *Service) AddNewEmployee(c *gin.Context) {
	method := postgresql.Service{}
	config := server.Service{}

	var empl postgresql.Employee

	err := c.BindJSON(&empl)

	if err != nil {
		fmt.Println("err: ", err)
	}

	result, error := config.LoadConfig("configs/server")

	if error != nil {
		fmt.Println("err: ", err)
	}

	method.AddNewEmployee(result, empl)
}

//AddNewOwner функция добавления новой организации
func (service *Service) AddNewOwner(c *gin.Context) {
	method := postgresql.Service{}
	config := server.Service{}

	var own postgresql.Owner

	err := c.BindJSON(&own)

	if err != nil {
		fmt.Println("err: ", err)
	}

	result, error := config.LoadConfig("configs/server")

	if error != nil {
		fmt.Println("err: ", err)
	}

	method.AddNewOwner(result, own)
}
