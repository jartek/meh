package models

import (
	"database/sql"
	"log"
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

func GetAllStadiums(db *sql.DB) []Stadium {
	rows, err := db.Query("SELECT * from stadiums")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()
	stadiums := []Stadium{}
	for rows.Next() {
		stadium := Stadium{}
		err := rows.Scan(&stadium.Id, &stadium.Name, &stadium.Location, &stadium.Capacity, &stadium.Altitude, &stadium.YearBuilt)
		if err != nil {
			log.Fatalln(err)
		}
		stadiums = append(stadiums, stadium)
	}

	return stadiums
}

func GetStadium(db *sql.DB, params martini.Params) Stadium {
	stadium := Stadium{}
	id, _ := strconv.Atoi(params["id"])
	row, err := db.Query("SELECT * FROM stadiums WHERE id = $1", id)
	if err != nil {
		log.Fatalln(err)
	}
	if row.Next() {
		row.Scan(&stadium.Id, &stadium.Name, &stadium.Location, &stadium.Capacity, &stadium.Altitude, &stadium.YearBuilt)
	}
	return stadium
}
