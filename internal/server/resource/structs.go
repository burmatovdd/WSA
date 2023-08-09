package resource

import (
	"database/sql"
	"time"
)

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserAuth struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Access   string `json:"access"`
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
	NoResolve int `json:"no_resolve"`
	NewWaf    int `json:"new_waf"`
}

type URL struct {
	Url   string `json:"url"`
	Email string `json:"email"`
	Owner string `json:"owner"`
}

type UrlTable struct {
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

type ResourceTable struct {
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
	URL     string `json:"URL"`
	Status  bool   `json:"Status"`
	WAF     bool   `json:"WAF"`
	SSL     bool   `json:"SSL"`
	DateEnd string `json:"DateEnd"`
	Email   string `json:"Email"`
}

type UpdateData struct {
	Url   string `json:"url"`
	Email string `json:"email"`
}
