package main

import (
	"os"

	"github.com/jartek/worldcup/models"
	"github.com/jartek/worldcup/server"
)

func main() {
	db := models.SetupDB(os.Getenv("DB_NAME"))
	server := worldcup.NewServer(db)
	server.Run()
}
