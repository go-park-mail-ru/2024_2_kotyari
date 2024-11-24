FROM golang:1.23.1-alpine3.20 AS build

WORKDIR /var/backend

# Copy necessary files
COPY cmd/rating_updater/main.go main.go
COPY . .

# Install dependencies and build
RUN go mod tidy
RUN go build -o rating_updater main.go

FROM alpine:edge as prod
RUN apk add bash
COPY --from=build /var/backend/rating_updater /app/rating_updater
COPY --from=build /var/backend/.env /app/.env

RUN mkdir -p /delivery/files
COPY files /app/files

RUN mkdir -p /app/internal/configs
COPY configs app/configs

WORKDIR /app
EXPOSE 8004
ENTRYPOINT ./rating_updater
