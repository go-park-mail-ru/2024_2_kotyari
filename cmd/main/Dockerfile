FROM golang:1.23.1-alpine3.20 AS build

WORKDIR /var/backend

# Копируем go.mod и go.sum для установки зависимостей
COPY cmd/main/main.go main.go
COPY . .

RUN go mod tidy
RUN go build -o main main.go

FROM alpine:edge as prod

RUN apk add bash

COPY --from=build /var/backend/main /app/main
COPY --from=build /var/backend/.env /app/.env

RUN mkdir -p /app/files
COPY files /app/files

WORKDIR /app
EXPOSE 8000

ENTRYPOINT ./main