package models

import (
	"database/sql"
	"strconv"

	"github.com/go-martini/martini"
)

type Stadium struct {
	Id        int
	Name      string
	Location  string
	Capacity  int
	Altitude  int
	YearBuilt int `db:"year_built"`
}

func GetAllStadiums(db *sql.DB) ([]interface{}, error) {
	return GetAll(db, &Stadium{})
}

func GetStadium(db *sql.DB, params martini.Params) (interface{}, error) {
	id, _ := strconv.Atoi(params["id"])
	return GetOne(db, &Stadium{}, id)
}
