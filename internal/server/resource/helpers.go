package resource

import (
	"WAF_Analytics/configs/serverConf"
	"WAF_Analytics/internal/server/postgresql/helpers"
	"encoding/json"
	"fmt"
	"log"
)

func checkDataInDB(query string, args []any) bool {
	rows, err := helpers.Select(query, args, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		return false
	}
	return true
}

func jsonParse(variable any) []byte {
	jsonStr, err := json.Marshal(variable)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	return jsonStr
}

func check(variable string) bool {
	if variable == "Error resolve and not curl" ||
		variable == "Not Waf" ||
		variable == "Error certificate" {
		return false
	} else {
		return true
	}
}

func getOwnerId(query string, args []any) int {
	if !checkDataInDB("select * from usdata where emailus = $1", args) {
		return 0
	}
	id := 0
	rows, err := helpers.Select(query, args, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		log.Fatalln("error: ", err)
		return 0
	}
	for rows.Next() {
		p := own{}
		err = rows.Scan(
			&p.ID,
			&p.NameOwn,
			&p.ShortName,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		id = p.ID
	}
	return id
}

func getUserId(query string, args []any) int {
	if !checkDataInDB("select * from usdata where emailus = $1", args) {
		return 0
	}
	id := 0
	rows, err := helpers.Select(query, args, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		log.Fatalln("error: ", err)
		return 0
	}

	for rows.Next() {
		p := user{}
		err = rows.Scan(
			&p.ID,
			&p.Email,
			&p.FIO,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		id = p.ID
	}
	return id
}

func getUserEmail(query string, args []any) string {
	email := ""
	rows, err := helpers.Select(query, args, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		log.Fatalln("error: ", err)
		return ""
	}
	for rows.Next() {
		p := user{}
		err = rows.Scan(
			&p.ID,
			&p.Email,
			&p.FIO,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		email = p.Email
	}
	return email
}

func counterUrl(query string, args []any) (urlNumber, error) {
	rows, err := helpers.Select(query, args, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		return urlNumber{}, err
	}
	k := 0
	req := urlNumber{}
	for rows.Next() {
		p := urlNumber{}
		err = rows.Scan(
			&p.Url,
			&p.Number,
		)
		if err != nil {
			return urlNumber{}, err
		}
		k++
		req = urlNumber{
			p.Url,
			k,
		}
	}
	return req, nil
}

func counterWaf(query string, args []any) (waf, error) {
	rows, err := helpers.Select(query, args, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		return waf{}, err
	}

	req := waf{}
	k := 0
	for rows.Next() {
		p := waf{}
		err = rows.Scan(
			&p.Waf,
			&p.Number,
		)
		if err != nil {
			return waf{}, err
		}
		if p.Waf != "Not Waf" {
			k++
		}
		req = waf{
			p.Waf,
			k}
	}
	return req, nil

}

func countUsers(query string, args []any) (userNumber, error) {
	rows, err := helpers.Select(query, args, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		return userNumber{}, err
	}
	req := userNumber{}
	k := 0
	for rows.Next() {
		p := userNumber{}
		err = rows.Scan(
			&p.IDOwner,
			&p.Number,
		)
		if err != nil {
			return userNumber{}, err
		}
		k++
		req = userNumber{
			p.IDOwner,
			k}
	}
	return req, nil
}
