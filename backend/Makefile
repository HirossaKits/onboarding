server:
	go run ./cmd/server/main.go

deamon:
	air

generate:
	go generate ./internal/db/ent

migrate:
	go run ./cmd/migrate/main.go

.PHONY: server deamon generate migrate