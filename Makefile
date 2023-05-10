DB_DSN="postgres://postgres:postgres@localhost:5432/storage?sslmode=disable"
MIGRATION_DIR=db/migrations

run:
	DB_DSN=${DB_DSN} go run ./cmd/store/main.go

lint:
	golangci-lint run

migrate_up:
	migrate -database ${DB_DSN} -path ${MIGRATION_DIR} up

migrate_down:
	migrate -database ${DB_DSN} -path ${MIGRATION_DIR} down

make test:
	go test -v ./...