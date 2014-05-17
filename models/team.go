package models

import (
	"database/sql"
	"log"
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
