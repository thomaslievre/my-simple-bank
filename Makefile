
# Simple Makefile for a Go project

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
	migrate -path db/migration -database "postgresql://bankroot:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://bankroot:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	@go run cmd/api/main.go

.PHONY: all build run test clean migrateup migratedown sqlc test server
		