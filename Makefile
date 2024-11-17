BINARY_NAME=server

default: help

include .env
export $(shell sed 's/=.*//' .env)

DB_URL := postgres://$(DB_USERNAME):$(DB_PASSWORD)@localhost:54320/$(DB_NAME)?sslmode=disable
MIGRATIONS_DIR := ./assets/migrations

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
	@echo 'apply-migrations - Apply all migrations from assets/migrations folder'
	@echo 'revert-migrations - Revert all migrations from assets/migrations folder'
	@echo 'For this tools to work you need to have migrate tool to be installed'
	@echo 'You can install it by running this command: go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest'

PROTO_DIR := ./api/protos

GEN_DIR := gen

PROTOC := protoc

ENTITIES := $(shell find $(PROTO_DIR) -mindepth 1 -maxdepth 1 -type d -exec basename {} \;)

proto-build: $(ENTITIES)


# export PATH=$PATH:$(go env GOPATH)/bin

$(ENTITIES):
	@echo "Генерация кода для сущности $@..."
	@mkdir -p $(PROTO_DIR)/$@/$(GEN_DIR)
	@$(PROTOC) \
		--proto_path=$(PROTO_DIR)/$@/proto \
		--go_out=$(PROTO_DIR)/$@/$(GEN_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_DIR)/$@/$(GEN_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/$@/proto/*.proto
	@echo "Генерация для $@ завершена."


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
	docker build -t main-go-image:latest .

all-run:
	docker compose up -d

main-refresh:
	docker stop main_go && docker rm main_go && docker rmi main-go-image && docker compose up -d

rating-updater-refresh:
	docker stop rating_updater_go && docker rm rating_updater_go && docker rmi rating-updater-go-image && docker compose up -d

pg-refresh:
	docker stop pg_db && docker rm pg_db && docker compose up -d

redis-refresh:
	docker stop redis_service && docker rm redis_service && docker compose up -d

all-stop:
	docker compose down

recreate-pg:
	docker compose down pg_db -v && docker compose up pg_db -d

recreate-redis:
	docker compose down redis_service -v && docker compose up redis_service -d

all-delete:
	docker compose down -v

profile-refresh:
	docker stop profile_go && docker rm profile_go && docker rmi profile-go-image

all-refresh: main-refresh pg-refresh redis-refresh

apply-migrations:
	@echo 'Applying migrations...'
	@migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) up

revert-migrations:
	@echo 'Reverting migrations...'
	@migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down


back-refresh: main-refresh rating-updater-refresh profile-refresh

# Правило генерации для каждой сущности
$(ENTITIES):
	  @echo "Генерация кода для сущности $@..."
	  @mkdir -p $(PROTO_DIR)/$@/$(GEN_DIR)
	  @$(PROTOC) \
		--proto_path=$(PROTO_DIR)/$@/proto \
		--go_out=$(PROTO_DIR)/$@/$(GEN_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_DIR)/$@/$(GEN_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/$@/proto/*.proto
	  @echo "Генерация для $@ завершена."



.PHONY: clean build
