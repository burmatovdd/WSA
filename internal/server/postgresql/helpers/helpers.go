package helpers

import (
	"WAF_Analytics/configs/serverConf"
	"database/sql"
	"fmt"
)

type cs struct {
	cs string
}

func getCs(conString serverConf.Config) cs {
	return cs{
		conString.POSTGRESQL_CONNSTRING,
	}
}

func Select(query string, arg []any, conString serverConf.Config) (*sql.Rows, error) {
	db, err := sql.Open("postgres", getCs(conString).cs)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query, arg...)
	if err != nil {
		return nil, err
	}
	return rows, err
}

func Exec(query string, arg []any, conString serverConf.Config) bool {
	db, err := sql.Open("postgres", getCs(conString).cs)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer db.Close()

	_, err = db.Exec(query, arg...)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
