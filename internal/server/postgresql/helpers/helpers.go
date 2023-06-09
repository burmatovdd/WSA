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
		log.Fatalln("error ", err)
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

//func Insert(table string, args []string, values postgresql.Owner, config string) {
//	db, err := sql.Open("postgres", config)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	defer db.Close()
//
//	query := "INSERT INTO " + table + " ("
//
//	for i := 0; i < len(args); i++ {
//		if i+1 == len(args) {
//			query = query + args[i] + ") VALUES ("
//			continue
//		}
//		query = query + args[i] + ","
//	}
//
//	for i := 1; i <= len(values); i++ {
//		if i == len(values) {
//			query = query + "$" + strconv.Itoa(i) + ")"
//			continue
//		}
//		query = query + "$" + strconv.Itoa(i) + ","
//	}
//
//	fmt.Println("query: ", query)
//
//	_, err = db.Exec(
//		query,
//		values)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//}
