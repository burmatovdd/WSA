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

func (service *PgService) Login(c *gin.Context) {
	var data Login
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		return
	}

	if !checkUserInDB("select * from users where emailus = $1 and passwordus = $2", []any{data.Login, data.Password}) {
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

func (service *PgService) GetStat(c *gin.Context) {
	rows, err := helpers.Select("select * from stat", nil, serverConf.DefaultConfig)
	defer rows.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	req := RequestStatistic{}

	for rows.Next() {
		p := ResponseStatistic{}
		err := rows.Scan(
			&p.ID,
			&p.Date,
			&p.AllServers,
			&p.ErrorServers,
			&p.WorkServers,
			&p.WithWaf,
			&p.Possible,
			&p.WafProcPossible,
			&p.WafProc,
			&p.WithKas,
			&p.WafAndKas,
			&p.WafAndKasProc,
			&p.AllCertificate,
			&p.OkCertificate)

		if err != nil {
			continue
		}
		req = RequestStatistic{
			p.AllServers,
			p.ErrorServers,
			p.WorkServers,
			p.WithWaf,
			p.AllCertificate,
			p.OkCertificate,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": string(toJson(req)),
	})
}

func (service *PgService) GetWeekStat(c *gin.Context) {
	days := findWeeks(time.Now())
	format := "2006-01-02"

	collections := WeeksStatistic{
		Last:    collector([]any{days.Last.Monday.Format(format), days.Last.Friday.Format(format)}),
		Current: collector([]any{days.Current.Monday.Format(format), days.Current.Friday.Format(format)}),
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": string(toJson(collections)),
	})
}

func (service *PgService) AddResource(c *gin.Context) {
	data := URL{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}
	result := checkData(data)
	if result.Result == false {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	collection := collectInfo(data.Url)
	res := helpers.Exec(
		"INSERT INTO url (nameurl,ip,err,waf,idusd,idowner,commonname,issuer,datecert) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)",
		[]any{
			collection.Resolve.NameUrl,
			collection.Resolve.Ip,
			collection.Resolve.Status,
			collection.Resolve.Waf,
			result.UserID,
			result.OwnerId,
			collection.Certificate.CommonName,
			collection.Certificate.Issuer,
			collection.Certificate.DateCert,
		},
		serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}
	res = helpers.Exec("INSERT INTO resource (nameurl,ipfirst,datefirst,datenores,status,wafdate,wafip) VALUES ($1,$2,$3,$4,$5,$6,$7)",
		[]any{
			data.Url,
			collection.Resolve.Ip,
			time.Now().Format("2006-01-02"),
			collection.Resolve.DateNoRes,
			collection.Resolve.Status,
			collection.Resolve.WafDate,
			collection.Resolve.WafIp,
		},
		serverConf.DefaultConfig,
	)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}

func (service *PgService) CheckResource(c *gin.Context) {
	var data string
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		return
	}

	if !checkResourceInDB([]any{data}) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": false,
		})
		return
	}

	rows, err := helpers.Select("select * from url where nameurl = $1", []any{data}, serverConf.DefaultConfig)
	defer rows.Close()

	if err != nil {
		fmt.Println("err")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	resourceStr := CheckResource{}

	for rows.Next() {
		p := UrlTable{}
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
		resourceStr = CheckResource{
			URL:     p.URL.String,
			IP:      p.IP.String,
			Status:  checker(p.Err.String),
			WAF:     checker(p.Waf.String),
			SSL:     getCertificate(p.URL.String),
			DateEnd: p.EndDate.String,
			Email:   getUserData("select * from usdata where idusd = $1", []any{p.IDUser}).Email.String,
			FIO:     getUserData("select * from usdata where idusd = $1", []any{p.IDUser}).FIO.String,
			Owner:   getOwnerData("select * from owners where shortname = $1", []any{p.IDOwner}).FullName.String,
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": string(toJson(resourceStr)),
	})
}

func (service *PgService) DeleteResource(c *gin.Context) {
	var data string
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		return
	}

	if !checkResourceInDB([]any{data}) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": false,
		})
		return
	}

	res := helpers.Exec("delete from url where nameurl = $1", []any{data}, serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": res,
		})
		return
	}
	res = helpers.Exec("delete from resource where nameurl = $1", []any{data}, serverConf.DefaultConfig)
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

func (service *PgService) UpdateResource(c *gin.Context) {
	var data UpdateData
	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println("err: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		return
	}

	user := getUserData("select * from usdata where emailus = $1", []any{data.Email})
	//if user.ID.Valid == false {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"code": http.StatusInternalServerError,
	//		"body": false,
	//	})
	//	return
	//}

	res := helpers.Exec("update url set idusd = $1 where nameurl = $2", []any{user.ID, data.Url}, serverConf.DefaultConfig)
	if !res {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"body": res,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}

func (service *PgService) GetGeneralStat(c *gin.Context) {
	var resources, deactiveRes, owners, waf int
	rows, err := helpers.Select("select count(*) from url", nil, serverConf.DefaultConfig)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&resources); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
			})
			return
		}
	}

	rows, err = helpers.Select("select count(*) from owners", nil, serverConf.DefaultConfig)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&owners); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
			})
			return
		}
	}

	rows, err = helpers.Select("select withwaf from stat where idstat = (select max(idstat) from stat );", nil, serverConf.DefaultConfig)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&waf); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
			})
			return
		}
	}

	rows, err = helpers.Select("select erservers from stat;", nil, serverConf.DefaultConfig)
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(&deactiveRes); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": string(toJson(GeneralStat{
			resources, deactiveRes, owners, waf,
		})),
	})

}

func (service *PgService) GetCertificates(c *gin.Context) {
	month := findMonth()
	rows, err := helpers.Select("select * from url where datecert between $1 and $2", []any{month.Current, month.Next}, serverConf.DefaultConfig)
	defer rows.Close()

	if err != nil {
		fmt.Println("err")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
		})
		return
	}

	certificates := []Certificate{}

	for rows.Next() {
		p := UrlTable{}
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
		certificates = append(certificates, Certificate{p.URL.String, p.EndDate.String[0:10]})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"body": string(toJson(sortCertificates(month, certificates))),
	})
}
