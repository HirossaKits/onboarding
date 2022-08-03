package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"go-chi-api/ent"
	"go-chi-api/ent/migrate"

	"go-chi-api/routers"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	// logging
	entOptions := []ent.Option{ent.Debug()}

	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed loading env file: %v", err)
	}

	// connect db
	mc := mysql.Config{
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "localhost" + ":" + os.Getenv("MYSQL_PORT"),
		DBName:               os.Getenv("MYSQL_DBNAME"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	client, err := ent.Open("mysql", mc.FormatDSN(), entOptions...)
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}

	defer client.Close()

	ctx := context.Background()

	// migrate db
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// start server
	router := chi.NewRouter()
	routers.UserRouter(router)
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("failed starting server: %v", err)
	}
}
