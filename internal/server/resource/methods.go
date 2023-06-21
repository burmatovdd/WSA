package resource

import (
	"WAF_Analytics/configs/serverConf"
	"WAF_Analytics/internal/server/postgresql/helpers"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
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

func checkDataInDB(query string) bool {
	rows, err := helpers.Select(query, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		return true
	}
	return false
}

//GetStat функция вывода общей статистики
func (service *PgService) GetStat(c *gin.Context) {
	rows, err := helpers.Select("select * from stat", serverConf.DefaultConfig)
	defer rows.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
	}

	req := reqStat{}

	for rows.Next() {
		p := resStat{}
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

		req = reqStat{
			p.allServers,
			p.errorServers,
			p.workServers,
			p.withWaf,
			p.allCertificate,
			p.okDateCertificate,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": string(jsonParse(req)),
	})
}

//GetResStat функция вывода данных по каждому ресурсу
func (service *PgService) GetResStat(c *gin.Context) {
	rows, err := helpers.Select("select * from url", serverConf.DefaultConfig)
	defer rows.Close()

	if err != nil {
		fmt.Println("err")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
	}

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
				Status:  check(p.Err.String),
				WAF:     check(p.Waf.String),
				SSL:     check(p.Issuer.String),
				DateEnd: p.EndDate.String,
			},
		}

		resourceArr = append(resourceArr, resourceStr)
	}

	arr := request{
		resourceArr,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": string(jsonParse(arr)),
	})

}

//AddOwner функция добавления организации
func (service *PgService) AddOwner(c *gin.Context) {
	var own owner
	err := c.BindJSON(&own)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
	}

	args := []any{own.FullName, own.ShortName}

	if checkDataInDB("select * from owners where nameown = '" + own.FullName + "'") {
		fmt.Println("already exist")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": "already exist",
		})
	}
	res := helpers.Exec("INSERT INTO owners (nameown, shortname) VALUES ($1,$2)", args, serverConf.DefaultConfig)

	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": res,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": res,
	})
}

//AddEmployee функция добавления нового пользователя
func (service *PgService) AddEmployee(c *gin.Context) {
	var emp employee
	err := c.BindJSON(&emp)
	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
	}

	if checkDataInDB("select * from usdata where emailus = '" + emp.Email + "'") {
		fmt.Println("already exist")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": "already exist",
		})
	}

	args := []any{emp.Email, emp.Password, emp.Access}

	res := helpers.Exec("INSERT INTO users (emailus, passwordus, accessus) VALUES ($1,$2,$3)", args, serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": res,
		})
	}

	args = []any{emp.Email, emp.Initials}

	res = helpers.Exec("INSERT INTO usdata (emailus, fio) VALUES ($1,$2)", args, serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": res,
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code": http.StatusBadRequest,
		"body": res,
	})
}

//AddResource добавление нового ресурса
func (service *PgService) AddResource(c *gin.Context) {
	var resource resource
	err := c.BindJSON(&resource)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
	}

	userId := getUserId("select * from usdata where emailus = '" + resource.Email + "'")
	ownId := getOwnerId("select * from owners where shortname = '" + resource.Owner + "'")

	args := []any{resource.Url, resource.Ip, userId, ownId}

	if checkDataInDB("select * from resource where nameurl = '" + resource.Url + "'") {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": "already exist",
		})
	}

	res := helpers.Exec("INSERT INTO url (nameurl, ip,idusd,idowner) VALUES ($1,$2,$3,$4)", args, serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": res,
		})
	}

	args = []any{resource.Url, resource.Ip}
	res = helpers.Exec("INSERT INTO resource (nameurl, ipfirst) VALUES ($1,$2)", args, serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": res,
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code": http.StatusBadRequest,
		"body": res,
	})
}

func getOwnerId(query string) int {
	id := 0
	rows, err := helpers.Select(query, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		log.Fatalln("error: ", err)
		return 0
	}
	for rows.Next() {
		p := own{}
		err = rows.Scan(
			&p.ID,
			&p.NameOwn,
			&p.ShortName,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		id = p.ID
	}
	return id
}

func getUserId(query string) int {
	id := 0
	rows, err := helpers.Select(query, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		log.Fatalln("error: ", err)
		return 0
	}

	for rows.Next() {
		p := user{}
		err = rows.Scan(
			&p.ID,
			&p.Email,
			&p.FIO,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		id = p.ID
	}
	return id
}

func getUserEmail(query string) string {
	email := ""
	rows, err := helpers.Select(query, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		log.Fatalln("error: ", err)
		return ""
	}
	for rows.Next() {
		p := user{}
		err = rows.Scan(
			&p.ID,
			&p.Email,
			&p.FIO,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		email = p.Email
	}
	return email
}

func (service *PgService) FindResourceByOwner(c *gin.Context) {
	var name ownName
	err := c.BindJSON(&name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
	}
	ownId := getOwnerId("select * from owners where shortname = '" + name.Name + "'")
	fmt.Println("ownId: ", ownId)

	res := resourceByOwner{}
	var resArr []resourceByOwner

	fmt.Println("select * from url where idowner = " + strconv.Itoa(ownId) + "'")
	rows, err := helpers.Select("select * from url where idowner = '"+strconv.Itoa(ownId)+"'", serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
	}

	for rows.Next() {
		p := resourceInfo{}
		err = rows.Scan(
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
			fmt.Println("error in scan; ", err)
			continue
		}

		res = resourceByOwner{
			Url:      p.URL,
			Error:    check(p.Err.String),
			Waf:      check(p.Waf.String),
			DateCert: p.EndDate.String,
			Email:    getUserEmail("select * from usdata where idusd = '" + strconv.FormatInt(p.IDUser.Int64, 10) + "'"),
		}

		resArr = append(resArr, res)
	}
	req := ResByOwnReq{
		name.Name,
		resArr,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": string(jsonParse(req)),
	})
}
