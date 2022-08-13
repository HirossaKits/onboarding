package db

import (
	"context"
	"log"

	"go-chi-api/ent/migrate"
)

// migrate db
func main() {

	c := NewClient()
	defer c.Close()

	ctx := context.Background()

	err := c.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
