package server

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jartek/worldcup/models"
	"github.com/martini-contrib/render"

	"github.com/go-martini/martini"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

type Server *martini.ClassicMartini

var m *martini.Martini

var DBConfig struct {
	SslMode, Name, User, Password string
}

func LoadDBConfig() {
	DBConfig.Name = os.Getenv("DB_NAME")
	DBConfig.SslMode = os.Getenv("DB_SSLMODE")
	DBConfig.User = os.Getenv("DB_USER")
	DBConfig.Password = os.Getenv("DB_PASSWORD")
}

func SetupDB() *sql.DB {
	LoadDBConfig()
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		DBConfig.User, DBConfig.Password, DBConfig.Name, DBConfig.SslMode))
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

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

func CollectionResponse(r render.Render, collection []interface{}, err error) {
	if err != nil {
		r.JSON(404, nil)
	} else {
		r.JSON(200, collection)
	}
}

func IndividualResponse(r render.Render, collection interface{}, err error) {
	if err != nil {
		r.JSON(404, nil)
	} else {
		r.JSON(200, collection)
	}
}

func GetAllTeams(r render.Render, db *sql.DB) {
	teams, err := models.GetAllTeams(db)
	CollectionResponse(r, teams, err)
}

func GetTeam(r render.Render, db *sql.DB, params martini.Params) {
	team, err := models.GetTeam(db, params)
	IndividualResponse(r, team, err)
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
