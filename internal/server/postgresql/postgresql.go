package postgresql

import (
	"WAF_Analytics/configs/server"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type stat struct {
	ID            int
	date          time.Time
	allServers    int
	errorServers  int
	workServers   int
	withWaf       int
	wafProc       float64
	withKas       int
	wafAndKas     int
	wafAndKasProc float64
}

type Service struct {
	service *Postgresql
}

type Postgresql interface {
	GetAllStat()
}

//connection функция коннекта к базе данных
func connection(conString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", conString)
	if err != nil {
		fmt.Println(err.Error())
	}
	return db, err
}

//GetAllStat функция вывода данных из базы данных
func (service *Service) GetAllStat(config server.Config, err error) {
	db, err := connection(config.POSTGRESQL_CONNSTRING)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select * from stat")
	if err != nil {
		log.Fatalln("error: ", err)
	}
	defer rows.Close()

	statMap := make(map[string]int)

	for rows.Next() {
		p := stat{}
		err := rows.Scan(
			&p.ID,
			&p.date,
			&p.allServers,
			&p.errorServers,
			&p.workServers,
			&p.withWaf,
			&p.wafProc,
			&p.withKas,
			&p.wafAndKas,
			&p.wafAndKasProc)

		if err != nil {
			fmt.Println(err)
			continue
		}

		statMap = map[string]int{
			"allServers":   p.allServers,
			"errorServers": p.errorServers,
			"workServers":  p.workServers,
			"withWaf":      p.withWaf,
		}
	}

	jsonStr, err := json.Marshal(statMap)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(jsonStr))
	}
}
