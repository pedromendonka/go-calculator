run:
	go run ./cmd/main.go

build:
	go build -o calculator ./cmd/main.go

deps:
	go mod download

dev:
	air

lint:
	golangci-lint run
