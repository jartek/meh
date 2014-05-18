package server

import (
	"database/sql"

	"github.com/jartek/worldcup/models"
	"github.com/martini-contrib/render"

	"github.com/go-martini/martini"
)

type Server *martini.ClassicMartini

var m *martini.Martini

func NewServer(db *sql.DB) Server {
	m := Server(martini.Classic())
	m.Use(render.Renderer())
	m.Map(db)
	m.Get("/teams", GetAllTeams)
	m.Get("/teams/:id", GetTeam)
	m.Get("/stadiums", GetAllStadiums)
	m.Get("/stadiums/:id", GetStadium)
	return m
}

func GetAllTeams(r render.Render, db *sql.DB) {
	r.JSON(200, models.GetAllTeams(db))
}

func GetTeam(r render.Render, db *sql.DB, params martini.Params) {
	team := models.GetTeam(db, params)
	if team.Id != 0 {
		r.JSON(200, team)
	} else {
		r.JSON(404, nil)
	}
}

func GetAllStadiums(r render.Render, db *sql.DB) {
	r.JSON(200, models.GetAllStadiums(db))
}

func GetStadium(r render.Render, db *sql.DB, params martini.Params) {
	team := models.GetStadium(db, params)
	if team.Id != 0 {
		r.JSON(200, team)
	} else {
		r.JSON(404, nil)
	}
}
