default: run

run:
	go run cmd/main.go

run-prod:
	GIN_MODE=release go run cmd/main.go

test:
	go test -coverprofile=c.out ./...
	go tool cover -func=c.out