run:
	go run ./cmd/store/main.go

lint:
	golangci-lint run

MIGRATION_DIR=./migrations
TEST_MIGRATION_DIR=./migrations/test_data
MIGRATION_DSN="host=localhost port=5432 dbname=storage user=postgres password=postgres sslmode=disable"

migrate_status:
	goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} status -v

migrate_up:
	goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} up -v

migrate_down:
	goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} down -v

test_migrate_up:
	goose -dir ${TEST_MIGRATION_DIR} postgres ${MIGRATION_DSN} up -v

test_migrate_down:
	goose -dir ${TEST_MIGRATION_DIR} postgres ${MIGRATION_DSN} down -v