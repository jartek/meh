package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/go-martini/martini"
)

type Team struct {
	Id       int
	Name     string
	NickName string `db:"nick_name"`
}

func GetAllTeams(db *sql.DB) []Team {
	// rows, err := db.Query("SELECT id, name, nick_name from teams")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer rows.Close()
	// teams := []Team{}
	// columns, _ := rows.Columns()
	// values := make([]interface{}, len(columns))
	// valuePtrs := make([]interface{}, len(columns))
	// ctr := 0
	// for rows.Next() {
	// 	for i, _ := range columns {
	// 		valuePtrs[i] = &values[i]
	// 	}
	// 	err := rows.Scan(valuePtrs...)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	for i, _ := range columns {
	// 		b, err := values[i].([]byte)

	// 		if err {
	// 			values[i] = string(b)
	// 		}
	// 	}
	// 	// teams = append(teams, values[ctr])
	// 	ctr = ctr + 1
	// }

	//
	teams := []Team{}
	fmt.Println(GetAll(db, &Team{}))
	return teams
}

func GetTeam(db *sql.DB, params martini.Params) Team {
	t := Team{}
	id, _ := strconv.Atoi(params["id"])
	row, err := db.Query("SELECT id, name, nick_name FROM teams WHERE id = $1", id)
	if err != nil {
		log.Fatalln(err)
	}
	if row.Next() {
		row.Scan(&t.Id, &t.Name, &t.NickName)
	}
	return t
}
