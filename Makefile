#!make
include .env
APP_NAME := entain-technical-task
DB_CONTAINER := entain-task-db
DOCKER_COMPOSE_FILE := docker-compose.yml

run:
	go run cmd/server/main.go

build:
	go build cmd/server/main.go

compose:
	docker-compose up -d

migrateup:
	migrate -path internal/db/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/db/migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

generate:
	sqlc generate

.PHONY: run build compose migrateup migratedown