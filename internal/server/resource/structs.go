package resource

import (
	"database/sql"
	"time"
)

type resStat struct {
	ID              int
	date            time.Time
	allServers      int
	errorServers    int
	workServers     int
	withWaf         int
	possible        float64
	wafProcPossible float64
	wafProc         float64
	withKas         int
	wafAndKas       int
	wafAndKasProc   float64
	allCertificate  int
	okCertificate   int
}

type reqStat struct {
	AllServers      int `json:"allServers"`
	ErrorServers    int `json:"errorServers"`
	WorkingServers  int `json:"workingServers"`
	WithWaf         int `json:"withWaf"`
	AllCertificates int `json:"allCertificates"`
	OkCertificates  int `json:"okCertificates"`
}

type resourceInfo struct {
	ID         int
	URL        string
	IP         string
	Err        sql.NullString
	Waf        sql.NullString
	IDUser     sql.NullInt64
	IDOwner    sql.NullInt64
	CommonName string
	Issuer     sql.NullString
	EndDate    sql.NullString
}

type request struct {
	Resource []requestBody `json:"resources"`
}

type requestBody struct {
	Resource resourceReq `json:"resource"`
}

type resourceReq struct {
	URL     string `json:"URL"`
	Status  bool   `json:"Status"`
	WAF     bool   `json:"WAF"`
	SSL     bool   `json:"SSL"`
	DateEnd string `json:"DateEnd"`
	Email   string `json:"Email"`
}

type employee struct {
	Initials string `json:"initials"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Access   bool   `json:"access"`
}

type owner struct {
	FullName  string `json:"fullName"`
	ShortName string `json:"shortName"`
}

type resource struct {
	Url   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
	Owner string `json:"owner,omitempty"`
}

type user struct {
	ID    int    `json:"ID"`
	Email string `json:"email"`
	FIO   string `json:"FIO"`
}

type own struct {
	ID        int    `json:"ID"`
	NameOwn   string `json:"email"`
	ShortName string `json:"shortName"`
}

type ownName struct {
	Name string `json:"name"`
}

type resByOwnReq struct {
	OwnerName string `json:"ownerName"`
	Resourses []resourceByOwner
}

type resourceByOwner struct {
	Url      string `json:"url"`
	Error    bool   `json:"error"`
	Waf      bool   `json:"waf"`
	DateCert string `json:"dateCert"`
	Email    string `json:"email"`
}

type urlNumber struct {
	Url    string `json:"url"`
	Number int    `json:"number"`
}

type userNumber struct {
	IDOwner int `json:"IDOwner"`
	Number  int `json:"Users"`
}

type waf struct {
	Waf    string `json:"waf"`
	Number int    `json:"count"`
}

type ownerInfoReq struct {
	ID   int       `json:"ID"`
	Info ownerInfo `json:"Info"`
}

type ownerInfo struct {
	Url   int `json:"Url"`
	Waf   int `json:"Waf"`
	Users int `json:"Users"`
}

type resourceName struct {
	Name string `json:"name"`
}

type updateResource struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type resourceBody struct {
	ID        int            `json:"ID"`
	NameURL   string         `json:"NameURL"`
	IpFirst   sql.NullString `json:"IpFirst"`
	IpNow     sql.NullString `json:"IpNow"`
	DateFirst sql.NullTime   `json:"DateFirst"`
	Status    sql.NullString `json:"Status"`
	DateNoRes sql.NullTime   `json:"DateNoRes"`
	WafDate   sql.NullTime   `json:"WafDate"`
	WafIp     sql.NullString `json:"WafIp"`
}

type report struct {
	LastWeek    lastWeek    `json:"lastWeek"`
	CurrentWeek currentWeek `json:"currentWeek"`
}

type currentWeek struct {
	NoResolve int `json:"noResolve"`
	NewWaf    int `json:"newWaf"`
}

type lastWeek struct {
	NoResolve int `json:"noResolve"`
	NewWaf    int `json:"newWaf"`
}

type login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type users struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Access   bool   `json:"access"`
}
