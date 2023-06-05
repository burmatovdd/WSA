package postgresql

import (
	"WAF_Analytics/configs/server"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

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

//AddNewEmployee функция добавления нового пользователя
func (service *Service) AddNewEmployee(config server.Config, data Employee) {
	db, err := connection(config.POSTGRESQL_CONNSTRING)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	defer db.Close()

	_, err = db.Exec(
		"INSERT INTO users (emailus, passwordus, accessus) VALUES ($1,$2,$3)",
		data.Email, data.Password, data.Access)
	if err != nil {
		fmt.Println(err.Error())
	}
}

//AddNewOwner функция добавления организации
func (service *Service) AddNewOwner(config server.Config, data Owner) {
	db, err := connection(config.POSTGRESQL_CONNSTRING)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	defer db.Close()

	_, err = db.Exec(
		"INSERT INTO owners (nameown, shortname) VALUES ($1,$2)",
		data.FullName, data.ShortName)
	if err != nil {
		fmt.Println(err.Error())
	}
}

//func (service *Service) AddNewResource(config server.Config, err error) {
//
//	res := struct {
//		Url      string
//		Ip       string
//		Employee string
//		Email    string
//		Owner    string
//	}{Url: "test.ru", Ip: "192.168.0.1", Employee: "John Doe", Email: "johndoe@mail.ru", Owner: "Test"}
//
//	db, err := connection(config.POSTGRESQL_CONNSTRING)
//	if err != nil {
//		fmt.Println("err: ", err.Error())
//	}
//	defer db.Close()
//
//	_, err = db.Exec("INSERT INTO usersbonds (user_id, bond_id, count) VALUES ($1,$2,$3)", userID, bondId, count)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//}
