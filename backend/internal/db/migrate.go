package db

import (
	"context"
	"log"

	"go-chi-api/internal/db/ent/migrate"
)

// migrate db
func Migrate() {

	store := GetInstance()
	defer store.Client.Close()

	ctx := context.Background()

	err := store.Client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
