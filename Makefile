.PHONY: build test lint clean install

BINARY_NAME=log-analyzer
VERSION=$(shell git describe --tags --always --dirty)

build:
	@echo "Building ${BINARY_NAME} ${VERSION}"
	@go build -ldflags="-X 'main.Version=${VERSION}'" -o ${BINARY_NAME} ./cmd/log-analyzer

test:
	@echo "Running tests..."
	@go test -v -cover ./...

test-coverage:
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

lint:
	@echo "Running linter..."
	@golangci-lint run ./...

run: build
	@./${BINARY_NAME} -file test.log -stats

install:
	@go install ./cmd/log-analyzer

clean:
	@echo "Cleaning..."
	@rm -f ${BINARY_NAME}
	@rm -f output.json output.csv
	@rm -f coverage.out coverage.html
	@go clean

docker-build:
	@docker build -t log-analyzer:latest .

help:
	@echo "Available targets:"
	@echo "  build         - Build the binary"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  lint          - Run linter"
	@echo "  run           - Build and run with test file"
	@echo "  install       - Install to GOPATH"
	@echo "  clean         - Clean build artifacts"
	@echo "  docker-build  - Build Docker image"