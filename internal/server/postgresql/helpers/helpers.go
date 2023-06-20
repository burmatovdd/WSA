package helpers

import (
	"WAF_Analytics/configs/serverConf"
	"database/sql"
	"fmt"
	"log"
)

type cs struct {
	cs string
}

func getCs(conString serverConf.Config) cs {
	return cs{
		conString.POSTGRESQL_CONNSTRING,
	}
}

func Select(query string, conString serverConf.Config) (*sql.Rows, error) {
	db, err := sql.Open("postgres", getCs(conString).cs)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalln("error in select ", err)
	}
	return rows, err
}

func Exec(query string, arg []any, conString serverConf.Config) bool {
	db, err := sql.Open("postgres", getCs(conString).cs)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	_, err = db.Exec(query, arg...)
	if err != nil {
		log.Fatalln("error ", err)
		return false
	}
	return true
}
