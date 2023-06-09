.PHONY: help build build-local up down logs ps test
.DEFAULT_GOAL := help

DOCKER_TAG := latest
build: ## Build the docker image
	docker build -t docker.io/library/gotodo:$(DOCKER_TAG) --target deploy .

build-dev: ## Build the docker image
	docker build -t docker.io/library/gotodo:$(DOCKER_TAG) --target dev .

build-local: ## Build the docker image locally
	docker compose build --no-cache

up: ## Start the docker containers
	docker compose up -d

down: ## Stop the docker containers
	docker compose down

restart: ## Restart the docker containers
	docker compose restart

logs: ## Show the logs
	docker compose logs -f

ps: ## Show the running containers
	docker compose ps

test: ## Run the tests
	go test -race -shuffle=on ./...

## database
migrate: ## Run the database migrations
	mysqldef site -uadmin -padmin -h127.0.0.1 -P3306 --file=./_tools/mysql/schema.sql