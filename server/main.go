package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"go-chi-api/ent"
	"go-chi-api/ent/migrate"

	"github.com/go-chi/chi/v5"
	"github.com/go-sql-driver/mysql"
)

func main() {

	// logging
	entOptions := []ent.Option{ent.Debug()}

	// connect db
	mc := mysql.Config{
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		Net:                  "tcp",
		Addr:                 "localhost" + ":" + "3306",
		DBName:               "db",
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
	UserRouter(&router)
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("failed starting server: %v", err)
	}
}
