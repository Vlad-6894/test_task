include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	@docker compose up -d test-task-postgres

env-down:
	@docker compose down test-task-postgres

postgres-cleanup:
	@read -p "Очистить pg_data? Опасность утери данных. [y/N]: " choice; \
	if [ "$$choice" = "y" ] || [ "$$choice" = "Y" ]; then \
		docker compose down test-task-postgres port-forwarder && \
		sudo rm -rf ${PROJECT_ROOT}/out/pg_data && \
		echo "Очищено"; \
	else \
		echo "Операция отменена"; \
	fi

logs-cleanup:
	@read -p "Очистить все логи? Опасность утери важных логов!. [y/N]" choice; \
	if [ "$$choice" = "y" ] || [ "$$choice" = "Y" ]; then \
		sudo rm -rf ${PROJECT_ROOT}/out/logs && \
		echo "Очищено"; \
	else \
		echo "Операция отменена"; \
	fi


create-migrate:
	@if [ -z "$(seq)" ]; then \
		echo "Нет параметра seq"; \
		exit 1; \
	fi;
	docker compose run --rm postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Нет параметра action"; \
		exit 1; \
	fi;
	@docker compose run --rm postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@test-task-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

start-port-forward:
	@docker compose up -d port-forwarder

close-port-forward:
	@docker compose down port-forwarder

run-application:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=test-task-postgres && \
	go mod tidy && \
	go run cmd/main.go

deploy-golang-app:
	@docker compose up -d --build test-task-golang-app

undeploy-golang-app:
	@docker compose down test-task-golang-app

ps:
	@docker compose ps

swagger-gen:
	@docker compose run --rm swagger \
		init \
		-g cmd/main.go \
		-o docs \
		--parseInternal \
		--parseDependency