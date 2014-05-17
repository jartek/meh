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

func LoadDBConfig(DbName string) {
	DBConfig.Name = DbName
	DBConfig.SslMode = os.Getenv("DB_SSLMODE")
	DBConfig.User = os.Getenv("DB_USER")
}

func SetupDB(DbName string) *sql.DB {
	LoadDBConfig(DbName) //Hacky way of switching dbname, must change it later
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=%s",
		DBConfig.User, DBConfig.Name, DBConfig.SslMode))
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
