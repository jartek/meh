package models

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/go-martini/martini"
)

type Team struct {
	Id       int
	Name     string
	NickName string `db:"nick_name"`
}

func GetAllTeams(db *sql.DB) []interface{} {
	return GetAll(db, &Team{})
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
