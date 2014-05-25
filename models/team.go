package models

import (
	"database/sql"
	"strconv"

	"github.com/go-martini/martini"
)

type Team struct {
	Id       int
	Name     string
	NickName string `db:"nick_name"`
}

func GetAllTeams(db *sql.DB) ([]interface{}, error) {
	return GetAll(db, &Team{})
}

func GetTeam(db *sql.DB, params martini.Params) (interface{}, error) {
	id, _ := strconv.Atoi(params["id"])
	return GetOne(db, &Team{}, id)
}
