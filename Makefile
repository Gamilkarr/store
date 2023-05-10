run:
	go run ./cmd/store/main.go

lint:
	golangci-lint run

MIGRATION_DIR=db/migrations
MIGRATION_DSN="postgres://postgres:postgres@localhost:5432/storage?sslmode=disable"

migrate_up:
	migrate -database ${MIGRATION_DSN} -path ${MIGRATION_DIR} up

migrate_down:
	migrate -database ${MIGRATION_DSN} -path ${MIGRATION_DIR} down

make test:
