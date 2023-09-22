package resource

import (
	"WAF_Analytics/configs/serverConf"
	"WAF_Analytics/internal/server/postgresql/helpers"
	"context"
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net"
	"strconv"
	"time"
)

func checkUserInDB(query string, args []any) bool {
	rows, err := helpers.Select(query, args, serverConf.DefaultConfig)
	defer rows.Close()

	us := UserAuth{}

	for rows.Next() {
		p := UserAuth{}
		err = rows.Scan(
			&p.ID,
			&p.Email,
			&p.Password,
			&p.Access,
		)

		if err != nil {
			fmt.Println(err)
			return false
		}
		us = UserAuth{
			p.ID,
			p.Email,
			p.Password,
			p.Access,
		}
	}
	if us.Email == "" {
		return false
	}
	return true
}

func checkResourceInDB(args []any) bool {
	rows, err := helpers.Select("select * from resource where nameurl = $1", args, serverConf.DefaultConfig)
	defer rows.Close()

	res := ResourceTable{}

	for rows.Next() {
		p := ResourceTable{}
		err = rows.Scan(
			&p.ID,
			&p.NameURL,
			&p.IpFirst,
			&p.IpNow,
			&p.DateFirst,
			&p.Status,
			&p.DateNoRes,
			&p.WafDate,
			&p.WafIp,
		)

		if err != nil {
			fmt.Println(err)
			continue
		}
		res = ResourceTable{
			p.ID,
			p.NameURL,
			p.IpFirst,
			p.IpNow,
			p.DateFirst,
			p.Status,
			p.DateNoRes,
			p.WafDate,
			p.WafIp,
		}
	}
	if res.ID.Int32 == 0 {
		return false
	}
	return true
}

func checkData(args URL) CheckDataResult {
	user := User{}
	owner := Owner{}
	if args.Email != "" {
		user = getUserData("select * from usdata where emailus = $1", []any{args.Email})
		if user.ID.Valid == false {
			return CheckDataResult{UserID: sql.NullInt32{}, OwnerId: sql.NullInt32{}, Result: false}
		}
	}
	if args.Owner != "" {
		owner = getOwnerData("select * from owners where shortname = $1", []any{args.Owner})
		if owner.ID.Valid == false {
			return CheckDataResult{UserID: sql.NullInt32{}, OwnerId: sql.NullInt32{}, Result: false}
		}
	}
	if checkResourceInDB([]any{args.Url}) == true {
		return CheckDataResult{UserID: sql.NullInt32{}, OwnerId: sql.NullInt32{}, Result: false}
	}
	if user.ID.Int32 != 0 && owner.ID.Int32 != 0 {
		return CheckDataResult{UserID: user.ID, OwnerId: owner.ID, Result: true}
	}
	if user.ID.Int32 != 0 && owner.ID.Int32 == 0 {
		return CheckDataResult{UserID: user.ID, OwnerId: sql.NullInt32{}, Result: true}
	}
	if user.ID.Int32 == 0 && owner.ID.Int32 != 0 {
		return CheckDataResult{UserID: sql.NullInt32{}, OwnerId: owner.ID, Result: true}
	}
	return CheckDataResult{UserID: sql.NullInt32{}, OwnerId: sql.NullInt32{}, Result: true}
}

func toJson(variable any) []byte {
	jsonStruct, err := json.Marshal(variable)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	return jsonStruct
}

func findWeeks(today time.Time) Weeks {
	mondayDurationDays := int(today.Weekday()) - 1
	fridayDurationDays := 5 - int(today.Weekday())

	return Weeks{
		Last: Week{
			Monday: today.Add(-(time.Duration(24*(mondayDurationDays+7)) * time.Hour)),
			Friday: today.Add(-(time.Duration(24*(7-fridayDurationDays)) * time.Hour)),
		},
		Current: Week{
			Monday: today.Add(-(time.Duration(24*mondayDurationDays) * time.Hour)),
			Friday: today.Add(time.Duration(24*fridayDurationDays) * time.Hour),
		},
	}
}

func findMonth() Months {
	today := time.Now()
	month := ""
	if int(today.Month()) < 10 {
		month = "0" + strconv.Itoa(int(today.Month()))
	}
	nextMonth, err := time.Parse("2006-01-02", strconv.Itoa(today.Year())+"-"+month+"-"+strconv.Itoa(31))

	if err != nil {
		nextMonth, err = time.Parse("2006-01-02", strconv.Itoa(today.Year())+"-"+month+"-"+strconv.Itoa(30))
		if err != nil {
			nextMonth, err = time.Parse("2006-01-02", strconv.Itoa(today.Year())+"-"+month+"-"+strconv.Itoa(29))
			if err != nil {
				nextMonth, err = time.Parse("2006-01-02", strconv.Itoa(today.Year())+"-"+month+"-"+strconv.Itoa(28))
			}
		}
	}

	return Months{
		strconv.Itoa(today.Year()) + "-" + month + "-" + "0" + strconv.Itoa(1),
		nextMonth.AddDate(0, 1, 0).Format("2006-01-02"),
	}

}

func collector(args []any) WeekStatistic {
	NoResolve, arrayNoResolve := counter("select * from resource where datenores between $1 and $2", args)
	NewWaf, arrayNewWaf := counter("select * from resource where wafdate between $1 and $2", args)
	return WeekStatistic{
		NoResolve:      NoResolve,
		NewWaf:         NewWaf,
		NoResResource:  arrayNoResolve,
		NewWafResource: arrayNewWaf,
	}
}

func counter(query string, args []any) (int, []WeekStatisticResource) {
	rows, err := helpers.Select(query, args, serverConf.DefaultConfig)
	defer rows.Close()
	resources := []WeekStatisticResource{}

	for rows.Next() {
		p := ResourceTable{}
		err = rows.Scan(
			&p.ID,
			&p.NameURL,
			&p.IpFirst,
			&p.IpNow,
			&p.DateFirst,
			&p.Status,
			&p.DateNoRes,
			&p.WafDate,
			&p.WafIp,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if p.DateNoRes.Valid == false {
			resources = append(resources, WeekStatisticResource{
				p.NameURL.String,
				p.WafDate.Time.Format("2006-01-02"),
			})
		} else {
			resources = append(resources, WeekStatisticResource{
				p.NameURL.String,
				p.DateNoRes.Time.Format("2006-01-02"),
			})
		}

	}
	return len(resources), resources
}

func getUserData(query string, args []any) User {
	rows, err := helpers.Select(query, args, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		log.Fatalln("error: ", err)
		return User{sql.NullInt32{}, sql.NullString{}, sql.NullString{}}
	}

	us := User{}
	for rows.Next() {
		p := User{}
		err = rows.Scan(
			&p.ID,
			&p.Email,
			&p.FIO,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		us = User{
			p.ID,
			p.Email,
			p.FIO,
		}
	}
	return us
}

func getOwnerData(query string, args []any) Owner {
	rows, err := helpers.Select(query, args, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		log.Fatalln("error: ", err)
		return Owner{sql.NullInt32{}, sql.NullString{}, sql.NullString{}}
	}

	owner := Owner{}
	for rows.Next() {
		p := Owner{}
		err = rows.Scan(
			&p.ID,
			&p.FullName,
			&p.ShortName,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		owner = Owner{
			p.ID,
			p.FullName,
			p.ShortName,
		}
	}
	return owner
}

func sortCertificates(month Months, data []Certificate) CertificateInfo {

	current := []Certificate{}
	next := []Certificate{}

	for i := 0; i < len(data); i++ {
		if string(data[i].Date[5])+string(data[i].Date[6]) == string(month.Current[5])+string(month.Current[6]) {
			current = append(current, data[i])
			continue
		}
		next = append(next, data[i])
	}
	return CertificateInfo{
		Current: current,
		Next:    next,
	}
}

func resolver(url string) string {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			// можно подставить любой другой DNS
			return d.DialContext(ctx, network, "8.8.8.8:53")
		},
	}
	ips, err := r.LookupHost(context.Background(), url)
	if err != nil {
		return "--------------------"
	}
	return ips[0]
}

func resolverCollector(url string) ResolveInfo {
	var status string
	var wafip, waf string
	var wafbool bool
	var errStatus bool
	var datenores, wafdate *time.Time

	ip := resolver(url)
	today, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	wafGroup := []string{"185.120.189.120", "185.120.189.211", "185.120.189.214"}
	datenores, wafdate, wafip, waf = nil, nil, "", ""

	for i := 0; i < len(wafGroup); i++ {
		if ip == wafGroup[i] {
			wafip = wafGroup[i]
			wafdate = &today
			waf = wafip[12:15]
			break
		}
	}

	if ip == "--------------------" {
		datenores = &today
		status = "Error resolve and not curl"
	}

	if waf == "Not Waf" || waf == "" || waf == "Error connect" {
		wafbool = false
	} else {
		wafbool = true
	}

	if status != "Ok" {
		errStatus = false
	} else {
		errStatus = true
	}

	return ResolveInfo{
		Ip:        ip,
		Status:    status,
		ErrStatus: errStatus,
		DateNoRes: datenores,
		WafDate:   wafdate,
		Waf:       waf,
		WafStatus: wafbool,
		WafIp:     &wafip,
		NameUrl:   url,
	}
}

func getCertificate(url string) UrlCertificate {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	var certStatus bool

	certificate := UrlCertificate{}

	conn, err := tls.Dial("tcp", url+":443", conf)
	if err != nil {
		log.Println("Error in Dial", err)
		return UrlCertificate{}
	}
	defer conn.Close()
	certs := conn.ConnectionState().PeerCertificates
	for _, cert := range certs {
		if cert.Issuer.CommonName == "--------------------" || cert.Issuer.CommonName == "" {
			certStatus = false
		} else {
			certStatus = true
		}
		certificate = UrlCertificate{
			CommonName: cert.Subject.CommonName,
			Issuer:     cert.Issuer.CommonName,
			DateCert:   cert.NotAfter.Format("2006-01-02"),
			CertStatus: certStatus,
		}
		break
	}
	return certificate
}

func collectInfo(url string) AddResourceCollection {
	return AddResourceCollection{
		Resolve:     resolverCollector(url),
		Certificate: getCertificate(url),
	}

}

func hashPassword(password string) []byte {
	pass := []byte(password)
	cost := 10

	hash, err := bcrypt.GenerateFromPassword(pass, cost)
	if err != nil {
		log.Fatalln("err: ", err)
	}

	return hash
}

func generateToken(login string, password string) (string, error) {
	var signingKey = "B0M9H%nWrF#wOhr6yKhn#h%5Db"

	token := jwt.NewWithClaims(jwt.SigningMethodES256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		login, password,
	})

	return token.SignedString([]byte(signingKey))
}
