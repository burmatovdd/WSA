package resource

import (
	"database/sql"
	"time"
)

type stat struct {
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

type resourceInfo struct {
	ID         int
	URL        string
	IP         string
	Err        sql.NullString
	Waf        sql.NullString
	IDUser     sql.NullInt64
	IDOwner    sql.NullInt64
	CommonName string
	Issuer     string
	EndDate    string
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

type ResByOwnReq struct {
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
