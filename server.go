package main

import (
	"database/sql"

	"github.com/jartek/worldcup/models"
	"github.com/martini-contrib/render"

	"github.com/go-martini/martini"
)

var m *martini.Martini

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Map(models.SetupDB())
	m.Get("/teams", GetAllTeams)
	m.Run()
}

func GetAllTeams(r render.Render, db *sql.DB) {
	r.JSON(200, models.GetAllTeams(db))
}
