# Simple Makefile for a Go project

DB_SOURCE=postgresql://bankroot:secret@postgres:5432/simple_bank?sslmode=disable

# Build the application
all: build

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Test the application
# @go test ./...
test:
	@echo "Testing..."
	go test -v -cover -short ./...

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

migrateup:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down 1

sqlc:
	sqlc generate

server:
	@go run cmd/api/main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/thomaslievre/my-simple-bank/db/sqlc Store

.PHONY: all build run test clean migrateup migratedown sqlc test server mock
	