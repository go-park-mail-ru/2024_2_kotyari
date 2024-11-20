FROM golang:1.23.1-alpine3.20 AS build

WORKDIR /var/backend

RUN apk update && apk add --no-cache git


# Copy necessary files
COPY cmd/main main
COPY internal internal
COPY .env .env
COPY docs docs

COPY go.mod go.mod
COPY go.sum go.sum

# Install dependencies and build
RUN go mod tidy
RUN go build -o main ./main

FROM alpine:edge as prod
RUN apk add bash
COPY --from=build /var/backend/main /app/main
COPY --from=build /var/backend/.env /app/.env

RUN mkdir -p /app/files
COPY files /app/files

WORKDIR /app
EXPOSE 8000
ENTRYPOINT ./main
