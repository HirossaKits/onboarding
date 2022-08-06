package db

import (
	"go-chi-api/ent"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func NewClient() *ent.Client {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed loading env file: %v", err)
	}

	mc := mysql.Config{
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("API_ROOT") + ":" + os.Getenv("MYSQL_PORT"),
		DBName:               os.Getenv("MYSQL_DBNAME"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	entOptions := []ent.Option{ent.Debug()}

	client, err := ent.Open("mysql", mc.FormatDSN(), entOptions...)
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}

	return client
}
