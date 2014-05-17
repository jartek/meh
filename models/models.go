package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var DBConfig struct {
	SslMode, Name, User string
}

func LoadDBConfig() {
	DBConfig.SslMode = os.Getenv("DB_SSLMODE")
	DBConfig.Name = os.Getenv("DB_NAME")
	DBConfig.User = os.Getenv("DB_USER")
}

func SetupDB() *sqlx.DB {
	LoadDBConfig()
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=%s",
		DBConfig.User, DBConfig.Name, DBConfig.SslMode))
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
