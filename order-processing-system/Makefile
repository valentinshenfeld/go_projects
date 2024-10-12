.PHONY: generate build run test clean

generate:
	@echo "Generating code..."
	go generate ./...

build:
	@echo "Building..."
	go build -o bin/api cmd/api/main.go

run:
	@echo "Running..."
	go run cmd/api/main.go

test:
	@echo "Running tests..."
	go test -v ./...

clean:
	@echo "Cleaning..."
	rm -rf bin

.DEFAULT_GOAL := build