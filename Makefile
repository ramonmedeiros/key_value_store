SHORT_HASH=$(shell git rev-parse --short HEAD)
IMAGE_NAME=key_store

default: run

build: bin/backend

bin/backend:
	go build -o bin/backend cmd/main.go

run:
	go run cmd/main.go

test:
	go test -coverprofile=c.out ./...
	go tool cover -func=c.out