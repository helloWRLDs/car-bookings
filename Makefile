#!make
include .env

set-path:
	export GOPATH=$$HOME/projects
	# export PATH=$$PATH:/usr/local/go/bin:$$GOPATH/bin
	export PATH=$$PATH:/usr/bin:/bin

run: 
	gnome-terminal --title="bookings" -- go run ./cmd/api

db-up:
	docker-compose up -d

db-down:
	docker-compose down

migrate-up: set-path db-up
	go install github.com/pressly/goose/v3/cmd/goose@v3.19.2
	goose -dir migrations ${DB} "postgresql://${PG_USER}:${PG_PASSWORD}@${PG_HOST}:${PG_PORT}/${PG_NAME}?sslmode=disable" up

migrate-down: set-path db-up
	goose -dir migrations ${DB} "postgresql://${PG_USER}:${PG_PASSWORD}@${PG_HOST}:${PG_PORT}/${PG_NAME}?sslmode=disable" down

migrate-up-to: set-path db-up
	goose -dir migrations ${DB} "postgresql://${PG_USER}:${PG_PASSWORD}@${PG_HOST}:${PG_PORT}/${PG_NAME}?sslmode=disable" up-to ${VOL}

migrate-down-to: set-path db-up
	goose -dir migrations ${DB} "postgresql://${PG_USER}:${PG_PASSWORD}@${PG_HOST}:${PG_PORT}/${PG_NAME}?sslmode=disable" down-to ${VOL}

start: db-up run