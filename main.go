package main

import (
	"github.com/jartek/worldcup/models"
	"github.com/jartek/worldcup/server"
)

func main() {
	db := models.SetupDB()
	server := server.NewServer(db)
	server.Run()
}
