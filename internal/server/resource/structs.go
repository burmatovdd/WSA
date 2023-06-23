package resource

import (
	"database/sql"
	"time"
)

type resStat struct {
	ID                int
	date              time.Time
	allServers        int
	errorServers      int
	workServers       int
	withWaf           int
	possible          float64
	wafProcPossible   float64
	wafProc           float64
	withKas           int
	wafAndKas         int
	wafAndKasProc     float64
	allCertificate    int
	okDateCertificate int
}

type reqStat struct {
	AllServers         int `json:"allServers"`
	ErrorServers       int `json:"errorServers"`
	WorkingServers     int `json:"workingServers"`
	WithWaf            int `json:"withWaf"`
	AllCertificates    int `json:"allCertificates"`
	OkDateCertificates int `json:"okDateCertificates"`
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
	Url      string `json:"url"`
	Ip       string `json:"ip"`
	Employee string `json:"employee"`
	Email    string `json:"email"`
	Owner    string `json:"owner"`
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
