include .env
export

MIGRATIONS_DIR=db/migrations
DRIVER=postgres
DB_DSN=postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_MODE)

migrate-up:
	goose -dir $(MIGRATIONS_DIR) $(DRIVER) "$(DB_DSN)" up

migrate-down:
	goose -dir $(MIGRATIONS_DIR) $(DRIVER) "$(DB_DSN)" down

migrate-status:
	goose -dir $(MIGRATIONS_DIR) $(DRIVER) "$(DB_DSN)" status

create-migration:
ifndef name
	$(error Please specify a migration name: make create-migration name=some_name)
endif
	goose -dir $(MIGRATIONS_DIR) create $(name) sql
