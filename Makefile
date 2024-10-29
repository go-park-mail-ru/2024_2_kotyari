BINARY_NAME=server

default: help

help:
	@echo 'usage: make [target]'
	@echo 'targets:'
	@echo 'build - Explicitly building a binary'
	@echo 'run - Running an app (with build and clean)'
	@echo 'clean - Explicitly running go clean and removing binary'
	@echo 'test - Running go test for all files'
	@echo 'test-coverage - Running go test for all files with coverage, showing total % in console, opening report in browser'
	@echo 'fmt - Running go fmt for all files'
	@echo 'go-build - Building Docker image for Go application'
	@echo 'all-run - Running Docker containers in detached mode'
	@echo 'back-refresh - Refreshing Go application container'
	@echo 'pg-refresh - Refreshing PostgreSQL database container'
	@echo 'all-refresh - Refreshing both Go application and PostgreSQL database containers'
	@echo 'all-stop - Stop all docker containers'

build:
	go build -o ${BINARY_NAME} ./cmd/main.go

run:
	go run ./cmd/main.go

clean:
	go clean
	rm ${BINARY_NAME}

test:
	go test ./...

test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

fmt:
	go fmt ./...

go-build:
	docker build -t back-go-image:latest .

all-run:
	docker compose up -d

back-refresh:
	docker stop back_go && docker rm back_go && docker rmi back-go-image && docker compose up -d

pg-refresh:
	docker stop pg_db && docker rm pg_db && docker compose up -d

redis-refresh:
	docker stop redis_service && docker rm redis_service && docker compose up -d

all-stop:
	docker compose down

pg-delete:
	docker compose down pg_db -v

redis-delete:
	docker compose down redis_service -v

all-delete:
	docker compose down -v

all-refresh: back-refresh pg-refresh redis-refresh

.PHONY: clean build
