# Variables
GOFILES=$(shell find . -type f -name '*.go' -not -path "./vendor/*")
BINARY_NAME=isopod_adventure

# Targets
all: deps format lint build test

deps:
	@echo "Installing dependencies..."
	@go mod tidy

format:
	@echo "Running gofmt..."
	@gofmt -s -w $(GOFILES)

lint:
	@echo "Running go vet..."
	@go vet
	@echo "Running golangci-lint..."
	@golangci-lint run .
	@staticcheck ./...

build:
	@echo "Compiling the binary..."
	@go build -o $(BINARY_NAME) main.go

test:
	@echo "Running tests..."
	@go test ./...

docs:
	@echo "Generating documentation..."
	@godoc -http=:6060

clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)

install_dev_tools:
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@go install golang.org/x/tools/cmd/godoc@latest

.PHONY: all format lint build test clean
