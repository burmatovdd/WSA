package resource

import (
	"WAF_Analytics/configs/serverConf"
	"WAF_Analytics/internal/server/postgresql/helpers"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

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
			&p.okCertificate)

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
			p.okCertificate,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": string(toJson(req)),
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
				URL:     p.URL.String,
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
		"body": string(toJson(arr)),
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

	if !checkOwner("select * from owners where nameown = $1", args) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": false,
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

	if !checkEmployee("select * from usdata where emailus = $1", args) {
		fmt.Println("already exist")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": false,
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
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": res,
	})
}

//AddResource добавление нового ресурса
func (service *PgService) AddResource(c *gin.Context) {
	var resource resource
	err := c.BindJSON(&resource)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	args := []any{resource.Email}
	userId := getUserId("select * from usdata where emailus = $1", args)

	args = []any{resource.Owner}
	ownId := getOwnerId("select * from owners where shortname = $1", args)

	args = []any{resource.Url}

	if !checkResourceInDB("select * from resource where nameurl = $1", args) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": false,
		})
		return
	}

	args = []any{resource.Url, userId, ownId}
	fmt.Println(args)
	res := helpers.Exec("INSERT INTO url (nameurl,idusd,idowner) VALUES ($1,$2,$3)", args, serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": res,
		})
		return
	}

	args = []any{resource.Url}
	res = helpers.Exec("INSERT INTO resource (nameurl) VALUES ($1)", args, serverConf.DefaultConfig)
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

//FindResourceByOwner поиск ресурсов по владельцу
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

	if !checkOwner("select * from owners where shortname = $1", args) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": false,
		})
		return
	}

	ownId := getOwnerId("select * from owners where shortname = $1", args)
	if ownId == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": false,
		})
		return
	}

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
			Url:      p.URL.String,
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
		"body": string(toJson(req)),
	})
}

//GetInformationAboutOwner получение информации о каждом владельце
func (service *PgService) GetInformationAboutOwner(c *gin.Context) {
	rows, err := helpers.Select("select * from owners", nil, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		return
	}

	req := []ownerInfoReq{}

	for rows.Next() {
		p := own{}
		err = rows.Scan(
			&p.ID,
			&p.NameOwn,
			&p.ShortName,
		)
		args := []any{p.ID}
		users, _ := countUsers("select idusd, count(*) from url where idowner = $1 group by idusd", args)

		waf, _ := counterWaf("select waf, count(*) from url where idowner = $1 group by waf", args)

		url, _ := counterUrl("select nameurl, count(*) from url where idowner = $1 group by nameurl", args)

		req = append(req, ownerInfoReq{
			p.ID,
			ownerInfo{
				url.Number,
				waf.Number,
				users.Number,
			},
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"resource": string(toJson(req)),
	})
}

//DeleteOwner удаление владельца
func (service *PgService) DeleteOwner(c *gin.Context) {
	var name ownName
	err := c.BindJSON(&name)
	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		return
	}

	args := []any{name.Name}

	if !checkOwner("select * from owners where nameown = $1", args) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": false,
		})
		return
	}

	args = []any{name.Name, 0}
	res := helpers.Exec("update url set idowner = $2 where nameurl = $1", args, serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": res,
		})
		return
	}

	args = []any{name.Name}
	res = helpers.Exec("delete from owners where nameown = $1", args, serverConf.DefaultConfig)
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

//DeleteResource удаление ресурса
func (service *PgService) DeleteResource(c *gin.Context) {
	var name resourceName
	err := c.BindJSON(&name)
	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		return
	}
	args := []any{name.Name}

	if checkResourceInDB("select * from resource where nameurl = $1", args) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": false,
		})
		return
	}

	res := helpers.Exec("delete from url where nameurl = $1", args, serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": res,
		})
		return
	}
	res = helpers.Exec("delete from resource where nameurl = $1", args, serverConf.DefaultConfig)
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

//UpdateResource обновление ресурса
func (service *PgService) UpdateResource(c *gin.Context) {
	var resource updateResource
	err := c.BindJSON(&resource)
	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		return
	}

	args := []any{resource.Email}

	userId := getUserId("select * from usdata where emailus = $1", args)

	if userId == false {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": false,
		})
		return
	}
	args = []any{resource.Name, userId}
	res := helpers.Exec("update url set idusd = $2 where nameurl = $1", args, serverConf.DefaultConfig)
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

func (service *PgService) GetWeekStat(c *gin.Context) {
	lastMonday, lastFriday := findLastWeek(time.Now())
	args := []any{lastMonday.Format("2006-01-02"), lastFriday.Format("2006-01-02")}
	noResolve, newWaf := counter(c, args)

	LastWeek := lastWeek{
		NoResolve: noResolve,
		NewWaf:    newWaf,
	}

	monday, friday := findCurrentWeek(time.Now())
	args = []any{monday.Format("2006-01-02"), friday.Format("2006-01-02")}
	noResolve, newWaf = counter(c, args)

	CurrentWeek := currentWeek{
		NoResolve: noResolve,
		NewWaf:    newWaf,
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": string(toJson(report{
			LastWeek,
			CurrentWeek,
		})),
	})
}

func (service *PgService) Login(c *gin.Context) {
	var login login
	err := c.BindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		return
	}

	rows, err := helpers.Select("select * from users", nil, serverConf.DefaultConfig)
	defer rows.Close()

	args := []any{login.Login, login.Password}

	if !checkLogin("select * from users where emailus = $1 and passwordus = $2", args) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": true,
	})
}

func (service *PgService) CheckResource(c *gin.Context) {
	var name string
	err := c.BindJSON(&name)
	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		return
	}
	args := []any{name}

	if checkResourceInDB("select * from resource where nameurl = $1", args) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": false,
		})
		return
	}

	rows, err := helpers.Select("select * from url where nameurl = $1", args, serverConf.DefaultConfig)
	defer rows.Close()

	if err != nil {
		fmt.Println("err")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	resourceStr := requestBody{}

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
				URL:     p.URL.String,
				Status:  check(p.Err.String),
				WAF:     check(p.Waf.String),
				SSL:     check(p.Issuer.String),
				DateEnd: p.EndDate.String,
				Email:   getUserEmail("select * from usdata where idusd = $1", []any{p.IDUser}),
			},
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": string(toJson(resourceStr)),
	})
}
