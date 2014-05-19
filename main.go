package main

import "github.com/jartek/worldcup/server"

func main() {
	db := server.SetupDB()
	server := server.NewServer(db)
	server.Run()
}
