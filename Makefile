include .env
export

export PROJECT_ROOT=$(shell pwd)
export USER_ID := $(shell id -u)
export GROUP_ID := $(shell id -g)

env-up:
	@USER_ID=$(USER_ID) GROUP_ID=$(GROUP_ID) docker compose up -d todoapp-postgres

env-down:
	@docker compose down todoapp-postgres

env-cleanup:
	@read -p "Cleanup all volume files? [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todoapp-postgres port-forwarder && \
		sudo rm -rf out/pgdata && \
		echo "Envirenment cleanup."; \
	else \
		echo "Environment canceled."; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Missing required parameter 'seq'."; \
		exit 1; \
	fi; \
	USER_ID=$(USER_ID) GROUP_ID=$(GROUP_ID) docker compose run --rm todoapp-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Missing required parameter 'action'."; \
		exit 1; \
	fi; \
	docker compose run --rm todoapp-postgres-migrate \
		-path /migrations \
			-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
			"${action}"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

todoapp-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run cmd/todoapp/main.go



