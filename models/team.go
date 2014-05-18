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

func GetAllTeams(db *sql.DB) []Team {
	rows, err := db.Query("SELECT name, nick_name from teams")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	teams := []Team{}
	for rows.Next() {
		team := Team{}
		err := rows.Scan(&team.Name, &team.NickName)
		if err != nil {
			log.Fatalln(err)
		}
		teams = append(teams, team)
	}

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
