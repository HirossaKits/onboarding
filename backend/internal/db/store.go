package db

import (
	"go-chi-api/internal/db/ent"
	"log"
	"os"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"entgo.io/ent/dialect/sql"
)

type store struct {
	Client *ent.Client
}

var once sync.Once
var instance *store

func open() {
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

	drv, err := sql.Open("mysql", mc.FormatDSN())
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}

	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	instance = &store{
		Client: ent.NewClient(ent.Driver(drv)),
	}
}

func GetInstance() *store {
	once.Do(open)
	return instance
}
