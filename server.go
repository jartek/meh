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
	m.Get("/", TestIt)
	m.Run()
}

func TestIt(r render.Render) {
	r.JSON(200, map[string]interface{}{"hello": "world"})
}
