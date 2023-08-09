package resource

import (
	"WAF_Analytics/configs/serverConf"
	"WAF_Analytics/internal/server/postgresql/helpers"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func checkUserInDB(query string, args []any) bool {
	rows, err := helpers.Select(query, args, serverConf.DefaultConfig)
	defer rows.Close()

	us := UserAuth{}

	for rows.Next() {
		p := UserAuth{}
		err = rows.Scan(
			&p.ID,
			&p.Email,
			&p.Password,
			&p.Access,
		)

		if err != nil {
			fmt.Println(err)
			return false
		}
		us = UserAuth{
			p.ID,
			p.Email,
			p.Password,
			p.Access,
		}
	}
	if us.Email == "" {
		return false
	}
	return true
}

func checkResourceInDB(args []any) bool {
	rows, err := helpers.Select("select * from resource where nameurl = $1", args, serverConf.DefaultConfig)
	defer rows.Close()

	res := UrlTable{}

	for rows.Next() {
		p := UrlTable{}
		err = rows.Scan(
			&p.ID,
			&p.NameURL,
			&p.IpFirst,
			&p.IpNow,
			&p.DateFirst,
			&p.Status,
			&p.DateNoRes,
			&p.WafDate,
			&p.WafIp,
		)

		if err != nil {
			fmt.Println(err)
			continue
		}
		res = UrlTable{
			p.ID,
			p.NameURL,
			p.IpFirst,
			p.IpNow,
			p.DateFirst,
			p.Status,
			p.DateNoRes,
			p.WafDate,
			p.WafIp,
		}
	}
	if res.ID.Int32 == 0 {
		return false
	}
	return true
}

func checkData(args URL) CheckDataResult {
	user := User{}
	owner := Owner{}
	if args.Email != "-" {
		user = getUserData([]any{args.Email})
		if user.ID.Valid == false {
			return CheckDataResult{UserID: sql.NullInt32{}, OwnerId: sql.NullInt32{}, Result: false}
		}
	}
	if args.Owner != "-" {
		owner = getOwnerData([]any{args.Owner})
		if owner.ID.Valid == false {
			return CheckDataResult{UserID: sql.NullInt32{}, OwnerId: sql.NullInt32{}, Result: false}
		}
	}
	if checkResourceInDB([]any{args.Url}) {
		return CheckDataResult{UserID: sql.NullInt32{}, OwnerId: sql.NullInt32{}, Result: false}
	}
	if user.ID.Int32 != 0 && owner.ID.Int32 != 0 {
		return CheckDataResult{UserID: user.ID, OwnerId: owner.ID, Result: true}
	}
	if user.ID.Int32 != 0 && owner.ID.Int32 == 0 {
		return CheckDataResult{UserID: user.ID, OwnerId: sql.NullInt32{}, Result: true}
	}
	if user.ID.Int32 == 0 && owner.ID.Int32 != 0 {
		return CheckDataResult{UserID: sql.NullInt32{}, OwnerId: owner.ID, Result: true}
	}
	return CheckDataResult{UserID: sql.NullInt32{}, OwnerId: sql.NullInt32{}, Result: true}
}

func checker(str string) bool {
	if str == "Error resolve and not curl" ||
		str == "Not Waf" ||
		str == "Error certificate" ||
		str == "" {
		return false
	} else {
		return true
	}
}

func toJson(variable any) []byte {
	jsonStruct, err := json.Marshal(variable)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	return jsonStruct
}

func findWeeks(today time.Time) Weeks {
	mondayDurationDays := int(today.Weekday()) - 1
	fridayDurationDays := 5 - int(today.Weekday())

	return Weeks{
		Last: Week{
			Monday: today.Add(-(time.Duration(24*(mondayDurationDays+7)) * time.Hour)),
			Friday: today.Add(-(time.Duration(24*(7-fridayDurationDays)) * time.Hour)),
		},
		Current: Week{
			Monday: today.Add(-(time.Duration(24*mondayDurationDays) * time.Hour)),
			Friday: today.Add(time.Duration(24*fridayDurationDays) * time.Hour),
		},
	}
}

func collector(args []any) WeekStatistic {
	return WeekStatistic{
		NoResolve: counter("select count(*) from resource where datenores between $1 and $2", args),
		NewWaf:    counter("select count(*) from resource where wafdate between $1 and $2", args),
	}
}

func counter(query string, args []any) int {
	rows, err := helpers.Select(query, args, serverConf.DefaultConfig)
	defer rows.Close()
	var number int

	for rows.Next() {
		if err = rows.Scan(&number); err != nil {
			log.Fatal(err)
		}
	}
	return number
}

func getUserData(args []any) User {
	rows, err := helpers.Select("select * from usdata where emailus = $1", args, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		log.Fatalln("error: ", err)
		return User{sql.NullInt32{}, sql.NullString{}, sql.NullString{}}
	}

	us := User{}
	for rows.Next() {
		p := User{}
		err = rows.Scan(
			&p.ID,
			&p.Email,
			&p.FIO,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		us = User{
			p.ID,
			p.Email,
			p.FIO,
		}
	}
	return us
}

func getOwnerData(args []any) Owner {
	rows, err := helpers.Select("select * from owners where shortname = $1", args, serverConf.DefaultConfig)
	defer rows.Close()
	if err != nil {
		log.Fatalln("error: ", err)
		return Owner{sql.NullInt32{}, sql.NullString{}, sql.NullString{}}
	}

	owner := Owner{}
	for rows.Next() {
		p := Owner{}
		err = rows.Scan(
			&p.ID,
			&p.FullName,
			&p.ShortName,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		owner = Owner{
			p.ID,
			p.FullName,
			p.ShortName,
		}
	}
	return owner
}
