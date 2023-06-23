package resource

import (
	"WAF_Analytics/configs/serverConf"
	"WAF_Analytics/internal/server/postgresql/helpers"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
)

//todo: убрать конкатенацию везде

//GetStat функция вывода общей статистики
func (service *PgService) GetStat(c *gin.Context) {
	rows, err := helpers.Select("select * from stat", nil, serverConf.DefaultConfig)
	defer rows.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
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
	rows, err := helpers.Select("select * from url", nil, serverConf.DefaultConfig)
	defer rows.Close()

	if err != nil {
		fmt.Println("err")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
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
		return
	}

	args := []any{own.FullName}

	if checkDataInDB("select * from owners where nameown = $1", args) {
		fmt.Println("already exist")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": "already exist",
		})
		return
	}
	args = []any{own.FullName, own.ShortName}
	res := helpers.Exec("INSERT INTO owners (nameown, shortname) VALUES ($1,$2)", args, serverConf.DefaultConfig)

	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": res,
		})
		return
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
		return
	}

	args := []any{emp.Email}

	if checkDataInDB("select * from usdata where emailus = $1", args) {
		fmt.Println("already exist")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": "already exist",
		})
		return
	}

	args = []any{emp.Email, emp.Password, emp.Access}

	res := helpers.Exec("INSERT INTO users (emailus, passwordus, accessus) VALUES ($1,$2,$3)", args, serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": res,
		})
		return
	}

	args = []any{emp.Email, emp.Initials}

	res = helpers.Exec("INSERT INTO usdata (emailus, fio) VALUES ($1,$2)", args, serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": res,
		})
		return
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
		return
	}

	args := []any{resource.Email}
	userId := getUserId("select * from usdata where emailus = $1", args)

	args = []any{resource.Owner}
	ownId := getOwnerId("select * from owners where shortname = $1", args)

	args = []any{resource.Url}

	if checkDataInDB("select * from resource where nameurl = $1", args) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": "already exist",
		})
		return
	}

	args = []any{resource.Url, resource.Ip, userId, ownId}
	res := helpers.Exec("INSERT INTO url (nameurl, ip,idusd,idowner) VALUES ($1,$2,$3,$4)", args, serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": res,
		})
		return
	}

	args = []any{resource.Url, resource.Ip}
	res = helpers.Exec("INSERT INTO resource (nameurl, ipfirst) VALUES ($1,$2)", args, serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": res,
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code": http.StatusBadRequest,
		"body": res,
	})
}

func (service *PgService) FindResourceByOwner(c *gin.Context) {
	var name ownName
	err := c.BindJSON(&name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		return
	}

	args := []any{name.Name}
	ch := checkDataInDB("select * from owners where shortname = $1", args)

	if !ch {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": "no data in data base",
		})
		return
	}

	ownId := getOwnerId("select * from owners where shortname = $1", args)

	res := resourceByOwner{}
	var resArr []resourceByOwner

	args = []any{ownId}
	rows, err := helpers.Select("select * from url where idowner = $1", args, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
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
			fmt.Println("error: ", err)
			continue
		}
		args = []any{p.IDUser.Int64}
		res = resourceByOwner{
			Url:      p.URL,
			Error:    check(p.Err.String),
			Waf:      check(p.Waf.String),
			DateCert: p.EndDate.String,
			Email:    getUserEmail("select * from usdata where idusd = $1", args),
		}

		resArr = append(resArr, res)
	}
	req := resByOwnReq{
		name.Name,
		resArr,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": string(jsonParse(req)),
	})
}

func (service *PgService) GetInformationAboutOwner(c *gin.Context) {
	rows, err := helpers.Select("select * from owners", nil, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		return
	}

	req := []ownerInfo{}

	for rows.Next() {
		p := own{}
		err = rows.Scan(
			&p.ID,
			&p.NameOwn,
			&p.ShortName,
		)
		if err != nil {
			continue
		}
		args := []any{p.ID}
		users, err := countUsers("select idusd, count(*) from url where idowner = $1 group by idusd", args)
		if err != nil {
			continue
		}
		waf, err := counterWaf("select waf, count(*) from url where idowner = $1 group by waf", args)
		if err != nil {
			continue
		}
		url, err := counterUrl("select nameurl, count(*) from url where idowner = $1 group by nameurl", args)
		if err != nil {
			continue
		}
		req = append(req, ownerInfo{
			p.ID,
			url.Number,
			waf.Number,
			users.Number,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"resource": string(jsonParse(req)),
	})
}
