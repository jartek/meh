package main

import (
	"database/sql"
	"os"

	"github.com/jartek/worldcup/models"
	"github.com/martini-contrib/render"

	"github.com/go-martini/martini"
)

var m *martini.Martini

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Map(models.SetupDB(os.Getenv("DB_NAME")))
	m.Get("/teams", GetAllTeams)
	m.Get("/teams/:id", GetTeam)
	m.Run()
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
