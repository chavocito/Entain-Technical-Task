APP_NAME := entain-tech-task
DOCKER_COMPOSE_FILE := docker-compose.yml

run:
	go run cmd/server/main.go

build:
	go build cmd/server/main.go

compose:
	docker-compose up -d

migrateup:
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

.PHONY: run build compose migrateup migratedown