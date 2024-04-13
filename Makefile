#!make
include .env

run:
	go run ./cmd/api

db-up:
	docker-compose up -d

db-down:
	docker-compose down

migrate-up: 
	@go install github.com/pressly/goose/v3/cmd/goose@latest 
	@goose -dir migrations ${DB} "postgresql://${PG_USER}:${PG_PASSWORD}@${PG_HOST}:${PG_PORT}/${PG_NAME}?sslmode=disable" up

migrate-down:
	@goose -dir migrations ${DB} "postgresql://${PG_USER}:${PG_PASSWORD}@${PG_HOST}:${PG_PORT}/${PG_NAME}?sslmode=disable" down

start: db-up run