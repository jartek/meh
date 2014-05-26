package models

import (
	"database/sql"
	"strconv"

	"github.com/go-martini/martini"
)

type Competition struct {
	Id        int64
	Name      string
	CreatedAt int64 `db:"created_at"`
	UpdatedAt int64 `db:"updated_at"`
}

func GetAllCompetitions(db *sql.DB) ([]interface{}, error) {
	return GetAll(db, &Competition{})
}

func GetCompetition(db *sql.DB, params martini.Params) (interface{}, error) {
	id, _ := strconv.Atoi(params["id"])
	return GetOne(db, &Competition{}, id)
}
