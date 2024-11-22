FROM golang:1.23.1-alpine3.20 AS build

WORKDIR /var/backend

# Copy necessary files
COPY cmd/main/main.go main.go
COPY internal internal
COPY .env .env
COPY docs docs
COPY api api

COPY go.mod go.mod
COPY go.sum go.sum

# Install dependencies and build
RUN go mod tidy
RUN go build -o main main.go

FROM alpine:edge as prod
RUN apk add bash
COPY --from=build /var/backend/main /app/main
COPY --from=build /var/backend/.env /app/.env

RUN mkdir -p /delivery/files
COPY files /app/files

WORKDIR /app
EXPOSE 8000
ENTRYPOINT ./main
