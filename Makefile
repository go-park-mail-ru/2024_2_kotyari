BINARY_NAME=server

default: help

help:
	@echo 'usage: make [target]'
	@echo 'targets:'
	@echo 'build - Explicitly building a binary'
	@echo 'run - Running an app (with build and clean)'
	@echo 'clean - Explicitly running go clean and removing binary'
	@echo 'test - Running go test for all files'
	@echo 'test-coverage: - Running go test for all files with coverage, showing total % in console, opening report in browser'
	@echo 'fmt - Running go fmt for all files'

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

docker-build:
	docker build -t back-go-image:latest .

docker-run:
	docker compose up -d

docker-refresh:
	docker stop back-go && docker rm back-go && docker rmi back-go-image && docker compose up -d

.PHONY: clean build
