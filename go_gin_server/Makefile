include .env

MIGRATION_PATH = ./db/migrations

.PHONY: migrate-create migrate-up migrate-down migrate-rollback

migrate-create:
	@migrate create -seq -ext sql -dir $(MIGRATION_PATH) $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_PATH) up

migrate-down:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_PATH) down

migrate-rollback:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_PATH) down 1


# migrate -database postgres://postgres:adminPassword@localhost:5432/tasks?sslmode=disable -path db/migrations up 1