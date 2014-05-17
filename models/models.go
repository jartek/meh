package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var DBConfig struct {
	SslMode, Name, User string
}

func LoadDBConfig() {
	DBConfig.Name = os.Getenv("DB_NAME")
	DBConfig.SslMode = os.Getenv("DB_SSLMODE")
	DBConfig.User = os.Getenv("DB_USER")
}

func SetupDB() *sql.DB {
	LoadDBConfig()
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=%s",
		DBConfig.User, DBConfig.Name, DBConfig.SslMode))
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
