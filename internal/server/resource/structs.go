package resource

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Login struct {
	Login    string `json:"login" db:"emailus"`
	Password string `json:"password"`
}

type TokenClaims struct {
	jwt.StandardClaims
	Login    string `json:"login"`
	Password string `json:"password"`
	Access   bool   `json:"access"`
}

type UserAuth struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Access   bool   `json:"access"`
}

type User struct {
	ID    sql.NullInt32  `json:"id"`
	Email sql.NullString `json:"email"`
	FIO   sql.NullString `json:"FIO"`
}

type RequestStatistic struct {
	AllServers      int `json:"allServers"`
	ErrorServers    int `json:"errorServers"`
	WorkingServers  int `json:"workingServers"`
	WithWaf         int `json:"withWaf"`
	AllCertificates int `json:"allCertificates"`
	OkCertificates  int `json:"okCertificates"`
}

type ResponseStatistic struct {
	ID              int
	Date            time.Time
	AllServers      int
	ErrorServers    int
	WorkServers     int
	WithWaf         int
	Possible        float64
	WafProcPossible float64
	WafProc         float64
	WithKas         int
	WafAndKas       int
	WafAndKasProc   float64
	AllCertificate  int
	OkCertificate   int
}

type Weeks struct {
	Last    Week
	Current Week
}

type Week struct {
	Monday time.Time
	Friday time.Time
}

type WeeksStatistic struct {
	Last    WeekStatistic `json:"last"`
	Current WeekStatistic `json:"current"`
}

type WeekStatistic struct {
	NoResolve      int                     `json:"no_resolve"`
	NewWaf         int                     `json:"new_waf"`
	NoResResource  []WeekStatisticResource `json:"no_res_resource"`
	NewWafResource []WeekStatisticResource `json:"new_waf_resource"`
}

type WeekStatisticResource struct {
	Resource string `json:"resource"`
	Date     string `json:"date"`
}

type URL struct {
	Url   string `json:"url"`
	Email string `json:"email"`
	Owner string `json:"owner"`
}

type ResourceTable struct {
	ID        sql.NullInt32  `json:"ID"`
	NameURL   sql.NullString `json:"NameURL"`
	IpFirst   sql.NullString `json:"IpFirst"`
	IpNow     sql.NullString `json:"IpNow"`
	DateFirst sql.NullTime   `json:"DateFirst"`
	Status    sql.NullString `json:"Status"`
	DateNoRes sql.NullTime   `json:"DateNoRes"`
	WafDate   sql.NullTime   `json:"WafDate"`
	WafIp     sql.NullString `json:"WafIp"`
}

type UrlTable struct {
	ID         sql.NullInt32
	URL        sql.NullString
	IP         sql.NullString
	Err        sql.NullString
	Waf        sql.NullString
	IDUser     sql.NullInt64
	IDOwner    sql.NullInt64
	CommonName sql.NullString
	Issuer     sql.NullString
	EndDate    sql.NullString
	ErrBool    sql.NullBool
	WafBool    sql.NullBool
	CertBool   sql.NullBool
}

type Owner struct {
	ID        sql.NullInt32  `json:"id"`
	FullName  sql.NullString `json:"nameOwner"`
	ShortName sql.NullString `json:"shortName"`
}

type CheckDataResult struct {
	UserID  sql.NullInt32
	OwnerId sql.NullInt32
	Result  bool
}

type CheckResource struct {
	URL       string         `json:"URL"`
	IP        string         `json:"IP"`
	Status    bool           `json:"Status"`
	WAF       bool           `json:"WAF"`
	SSLStatus bool           `json:"SSLStatus"`
	SSL       UrlCertificate `json:"SSL"`
	DateEnd   string         `json:"DateEnd"`
	Email     string         `json:"Email"`
	FIO       string         `json:"FIO"`
	Owner     string         `json:"Owner"`
}

type UpdateData struct {
	Url   string `json:"url"`
	Email string `json:"email"`
}

type GeneralStat struct {
	Resources          int `json:"resources"`
	DeactivateResource int `json:"deactivateResource"`
	Owners             int `json:"owners"`
	Waf                int `json:"waf"`
}

type CertificateInfo struct {
	Current []Certificate `json:"current"`
	Next    []Certificate `json:"next"`
}

type Certificate struct {
	Resource string `json:"resource"`
	Date     string `json:"date"`
}

type Months struct {
	Current string
	Next    string
}

type ResolveInfo struct {
	Ip        string     `json:"ip"`
	Status    string     `json:"status"`
	ErrStatus bool       `json:"errStatus"`
	DateNoRes *time.Time `json:"dateNoRes"`
	WafDate   *time.Time `json:"wafDate"`
	Waf       string     `json:"waf"`
	WafIp     *string    `json:"wafIp"`
	NameUrl   string     `json:"nameurl"`
	WafStatus bool       `json:"wafStatus"`
}

type UrlCertificate struct {
	CommonName string `json:"common_name"`
	Issuer     string `json:"issuer"`
	DateCert   string `json:"date_cert"`
	CertStatus bool   `json:"certStatus"`
}

type AddResourceCollection struct {
	Resolve     ResolveInfo    `json:"resolve"`
	Certificate UrlCertificate `json:"certificate"`
}

type AllStats struct {
	//GenStats GeneralStat `json:"genStats"`
	AllURL []Resource `json:"allURL"`
	//Owners []string `json:"owners"`
	WafURL []Resource `json:"wafURL"`
	ErrURL []Resource `json:"errURL"`
}

type Resource struct {
	Resource string `json:"resource"`
}

type SQLChart struct {
	Date       sql.NullString `json:"date"`
	AllServers sql.NullString `json:"allServers"`
	ErServers  sql.NullString `json:"erServers"`
	WithWAF    sql.NullString `json:"withWAF"`
}

type Month struct {
	Month string `json:"month"`
	Chart Chart  `json:"chart"`
}

type Chart struct {
	AllServers string `json:"allServers"`
	ErServers  string `json:"erServers"`
	WithWAF    string `json:"withWAF"`
}

type WAFStats struct {
	WithWAF int `json:"withWAF"`
	NoWAF   int `json:"noWAF"`
}
