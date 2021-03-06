package models

import (
	"database/sql"
	"strconv"

	"github.com/go-martini/martini"
)

type Stadium struct {
	Id        int64
	Name      string
	Location  string
	Capacity  int64
	Altitude  int64
	YearBuilt int64 `db:"year_built"`
	CreatedAt int64 `db:"created_at"`
	UpdatedAt int64 `db:"updated_at"`
}

func GetAllStadiums(db *sql.DB) ([]interface{}, error) {
	return GetAll(db, &Stadium{})
}

func GetStadium(db *sql.DB, params martini.Params) (interface{}, error) {
	id, _ := strconv.Atoi(params["id"])
	return GetOne(db, &Stadium{}, id)
}
