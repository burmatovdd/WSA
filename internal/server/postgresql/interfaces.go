package postgresql

import (
	"WAF_Analytics/configs/server"
)

type Service struct {
	service *Postgresql
}

type Postgresql interface {
	GetAllStat()
	GetResourcesStat(config server.Config, err error)
	AddNewResource(config server.Config, err error)
	AddNewEmployee(config server.Config, data Employee)
	AddNewOwner(config server.Config, data Employee)
}
