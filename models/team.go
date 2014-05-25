package models

import (
	"database/sql"
	"strconv"

	"github.com/go-martini/martini"
)

type Team struct {
	Id        int64
	Name      string
	NickName  string `db:"nick_name"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

func GetAllTeams(db *sql.DB) ([]interface{}, error) {
	return GetAll(db, &Team{})
}

func GetTeam(db *sql.DB, params martini.Params) (interface{}, error) {
	id, _ := strconv.Atoi(params["id"])
	return GetOne(db, &Team{}, id)
}
