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
	ID                int
	date              time.Time
	allServers        int
	errorServers      int
	workServers       int
	withWaf           int
	possible          float64
	wafProcPossible   float64
	wafProc           float64
	withKas           int
	wafAndKas         int
	wafAndKasProc     float64
	allCertificate    int
	okDateCertificate int
}

type resource struct {
	ID         int
	URL        string
	IP         string
	Err        string
	Waf        string
	IDUser     sql.NullString
	IDOwner    sql.NullString
	CommonName string
	Issuer     string
	EndDate    string
}

type request struct {
	Resource []requestBody `json:"resources"`
}

type requestBody struct {
	Recourse resourceReq `json:"recourse"`
}

type resourceReq struct {
	URL     string `json:"URL"`
	Status  bool   `json:"Status"`
	WAF     bool   `json:"WAF"`
	SSL     bool   `json:"SSL"`
	DateEnd string `json:"DateEnd"`
}

type Service struct {
	service *Postgresql
}

type Postgresql interface {
	GetAllStat()
	GetResourcesStat(config server.Config, err error)
}

//connection функция коннекта к базе данных
func connection(conString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", conString)
	if err != nil {
		fmt.Println(err.Error())
	}
	return db, err
}

func jsonParse(variable any) []byte {
	jsonStr, err := json.Marshal(variable)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	return jsonStr
}

func check(variable string) bool {
	if variable == "Error resolve and not curl" ||
		variable == "Not Waf" ||
		variable == "Error certificate" {
		return false
	} else {
		return true
	}
}

//GetAllStat функция вывода общей статистики
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
			&p.possible,
			&p.wafProcPossible,
			&p.wafProc,
			&p.withKas,
			&p.wafAndKas,
			&p.wafAndKasProc,
			&p.allCertificate,
			&p.okDateCertificate)

		if err != nil {
			fmt.Println(err)
			continue
		}

		statMap = map[string]int{
			"allServers":        p.allServers,
			"errorServers":      p.errorServers,
			"workServers":       p.workServers,
			"withWaf":           p.withWaf,
			"allCertificate":    p.allCertificate,
			"okDateCertificate": p.okDateCertificate,
		}
	}

	fmt.Println(string(jsonParse(statMap)))
}

//GetResourcesStat функция вывода данных по каждому ресурсу
func (service *Service) GetResourcesStat(config server.Config, err error) {
	db, err := connection(config.POSTGRESQL_CONNSTRING)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select * from url")
	if err != nil {
		log.Fatalln("error: ", err)
	}
	defer rows.Close()

	resourceStr := requestBody{}
	var resourceArr []requestBody

	for rows.Next() {
		p := resource{}
		err := rows.Scan(
			&p.ID,
			&p.URL,
			&p.IP,
			&p.Err,
			&p.Waf,
			&p.IDUser,
			&p.IDOwner,
			&p.CommonName,
			&p.Issuer,
			&p.EndDate)

		if err != nil {
			fmt.Println(err)
			continue
		}

		resourceStr = requestBody{
			resourceReq{
				URL:     p.URL,
				Status:  check(p.Err),
				WAF:     check(p.Waf),
				SSL:     check(p.Issuer),
				DateEnd: p.EndDate,
			},
		}

		resourceArr = append(resourceArr, resourceStr)
	}

	arr := request{
		resourceArr,
	}

	for i := 0; i < len(resourceArr); i++ {
		fmt.Println(resourceArr[i])
	}

	fmt.Println("-----------------------")

	fmt.Println(string(jsonParse(arr)))
}
