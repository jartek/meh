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
	err := db.QueryRow("SELECT id, name, nick_name FROM teams WHERE id = $1", id).Scan(&t.Id, &t.Name, &t.NickName)
	if err != nil {
		log.Fatalln(err)
	}
	return t
}
