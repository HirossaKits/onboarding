package main

import (
	"go-chi-api/internal/db"
)

func main() {
	db.Migrate()
}
