package main

import (
	"github.com/jartek/worldcup/models"
	"github.com/martini-contrib/render"

	"github.com/go-martini/martini"
)

var m *martini.Martini

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Map(models.SetupDB())
	m.Run()
}
