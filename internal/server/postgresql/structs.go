package postgresql

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

type resource struct {
	ID         int
	URL        string
	IP         string
	Err        string
	Waf        string
	IDUser     sql.NullString
	IDOwner    sql.NullString
	CommonName string
	Issuer     string
	EndDate    string
}

type request struct {
	Resource []requestBody `json:"resources"`
}

type requestBody struct {
	Recourse resourceReq `json:"recourse"`
}

type resourceReq struct {
	URL     string `json:"URL"`
	Status  bool   `json:"Status"`
	WAF     bool   `json:"WAF"`
	SSL     bool   `json:"SSL"`
	DateEnd string `json:"DateEnd"`
}

type Employee struct {
	Initials string `json:"initials"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Access   bool   `json:"access"`
}

type Owner struct {
	FullName  string `json:"fullName"`
	ShortName string `json:"shortName"`
}
