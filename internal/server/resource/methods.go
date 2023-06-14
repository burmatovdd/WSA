package resource

import (
	"WAF_Analytics/configs/serverConf"
	"WAF_Analytics/internal/server/postgresql/helpers"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

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

//GetStat функция вывода общей статистики
func (service *PgService) GetStat(c *gin.Context) {
	rows, _ := helpers.Select("select * from stat", serverConf.DefaultConfig)
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

//GetResStat функция вывода данных по каждому ресурсу
func (service *PgService) GetResStat(c *gin.Context) {
	rows, _ := helpers.Select("select * from url", serverConf.DefaultConfig)
	defer rows.Close()

	resourceStr := requestBody{}
	var resourceArr []requestBody

	for rows.Next() {
		p := resourceInfo{}
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

//AddOwner функция добавления организации
func (service *PgService) AddOwner(c *gin.Context) {
	var own owner
	err := c.BindJSON(&own)
	if err != nil {
		fmt.Println("err: ", err)
	}

	args := []any{own.FullName, own.ShortName}

	if checkDataInDB("select * from owners where nameown = '" + own.FullName + "'") {
		fmt.Println("already exist")
		return
	}
	res := helpers.Exec("INSERT INTO owners (nameown, shortname) VALUES ($1,$2)", args, serverConf.DefaultConfig)

	if !res {
		fmt.Println("error: ", res)
	}
	fmt.Println("res: ", res)
}

//AddEmployee функция добавления нового пользователя
func (service *PgService) AddEmployee(c *gin.Context) {
	var emp employee
	err := c.BindJSON(&emp)
	if err != nil {
		fmt.Println("err: ", err)
	}

	if checkDataInDB("select * from usdata where emailus = '" + emp.Email + "'") {
		fmt.Println("already exist")
		return
	}

	args := []any{emp.Email, emp.Password, emp.Access}

	res := helpers.Exec("INSERT INTO users (emailus, passwordus, accessus) VALUES ($1,$2,$3)", args, serverConf.DefaultConfig)
	if !res {
		fmt.Println("error: ", res)
	}
	fmt.Println("res: ", res)

	args = []any{emp.Email, emp.Initials}

	res = helpers.Exec("INSERT INTO usdata (emailus, fio) VALUES ($1,$2)", args, serverConf.DefaultConfig)
	if !res {
		fmt.Println("error: ", res)
	}
	fmt.Println("res: ", res)
}

//AddResource добавление нового ресурса
func (service *PgService) AddResource(c *gin.Context) {
	var resource resource
	err := c.BindJSON(&resource)
	if err != nil {
		fmt.Println("err: ", err)
	}

	args := []any{resource.Url, resource.Ip}

	if checkDataInDB("select * from resource where nameurl = '" + resource.Url + "'") {
		fmt.Println("already exist")
		return
	}

	res := helpers.Exec("INSERT INTO url (nameurl, ip) VALUES ($1,$2)", args, serverConf.DefaultConfig)
	if !res {
		fmt.Println("error: ", res)
	}
	fmt.Println("res: ", res)

	res = helpers.Exec("INSERT INTO resource (nameurl, ipfirst) VALUES ($1,$2)", args, serverConf.DefaultConfig)
	if !res {
		fmt.Println("error: ", res)
	}
	fmt.Println("res: ", res)
}

func checkDataInDB(query string) bool {
	rows, err := helpers.Select(query, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		return false
	}
	return true
}
