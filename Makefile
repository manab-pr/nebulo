# Nebulo Makefile

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt

# Binary names
SERVER_BINARY=nebulo-server
DEVICE_BINARY=nebulo-device

# Build directory
BUILD_DIR=bin

# Version
VERSION?=$(shell git describe --tags --always --dirty)
LDFLAGS=-ldflags "-w -s -X main.version=$(VERSION)"

.PHONY: help build build-server build-device clean test test-coverage run run-server run-device docker docker-build docker-run lint format tidy deps install dev

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: build-server build-device ## Build both server and device binaries

build-server: ## Build the main server binary
	@echo "Building main server..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(SERVER_BINARY) cmd/server/main.go

build-device: ## Build the device server binary
	@echo "Building device server..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(DEVICE_BINARY) cmd/device-server/main.go

clean: ## Clean build artifacts
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -rf dist
	rm -f coverage.out
	rm -f coverage.html

test: ## Run tests
	@echo "Running tests..."
	$(GOTEST) -v -race ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	$(GOTEST) -v -race -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

run: ## Run the main server (requires .env file)
	@echo "Starting main server..."
	$(GOCMD) run cmd/server/main.go

run-server: ## Run the main server (alias for run)
	@make run

run-device: ## Run the device server (requires .env file)
	@echo "Starting device server..."
	$(GOCMD) run cmd/device-server/main.go

docker: docker-build ## Build and run with Docker Compose
	docker-compose up -d

docker-build: ## Build Docker images
	@echo "Building Docker images..."
	docker build -t nebulo-server -f Dockerfile .
	docker build -t nebulo-device -f Dockerfile.device .

docker-run: ## Run with Docker Compose
	@echo "Starting services with Docker Compose..."
	docker-compose up -d

docker-stop: ## Stop Docker Compose services
	@echo "Stopping Docker Compose services..."
	docker-compose down

docker-logs: ## Show Docker Compose logs
	docker-compose logs -f

lint: ## Run golangci-lint
	@echo "Running linter..."
	golangci-lint run

format: ## Format Go code
	@echo "Formatting code..."
	$(GOFMT) ./...
	goimports -w .

tidy: ## Tidy Go modules
	@echo "Tidying modules..."
	$(GOMOD) tidy
	$(GOMOD) verify

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	$(GOMOD) download

install: build ## Install binaries to GOPATH/bin
	@echo "Installing binaries..."
	cp $(BUILD_DIR)/$(SERVER_BINARY) $(GOPATH)/bin/
	cp $(BUILD_DIR)/$(DEVICE_BINARY) $(GOPATH)/bin/

dev: ## Set up development environment
	@echo "Setting up development environment..."
	@if [ ! -f .env ]; then \
		echo "Copying .env.example to .env..."; \
		cp .env.example .env; \
	fi
	@echo "Installing development tools..."
	$(GOGET) -u github.com/golangci/golangci-lint/cmd/golangci-lint
	$(GOGET) -u golang.org/x/tools/cmd/goimports
	@echo "Development environment ready!"
	@echo "Edit .env file with your configuration, then run 'make run' to start the server."

# Release targets
release-build: ## Build release binaries for multiple platforms
	@echo "Building release binaries..."
	@mkdir -p dist
	# Linux AMD64
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o dist/$(SERVER_BINARY)-linux-amd64 cmd/server/main.go
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o dist/$(DEVICE_BINARY)-linux-amd64 cmd/device-server/main.go
	# Linux ARM64
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o dist/$(SERVER_BINARY)-linux-arm64 cmd/server/main.go
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o dist/$(DEVICE_BINARY)-linux-arm64 cmd/device-server/main.go
	# macOS AMD64
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o dist/$(SERVER_BINARY)-darwin-amd64 cmd/server/main.go
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o dist/$(DEVICE_BINARY)-darwin-amd64 cmd/device-server/main.go
	# macOS ARM64
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o dist/$(SERVER_BINARY)-darwin-arm64 cmd/server/main.go
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o dist/$(DEVICE_BINARY)-darwin-arm64 cmd/device-server/main.go
	# Windows AMD64
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o dist/$(SERVER_BINARY)-windows-amd64.exe cmd/server/main.go
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o dist/$(DEVICE_BINARY)-windows-amd64.exe cmd/device-server/main.go

# Database management
db-up: ## Start MongoDB and Redis with Docker
	docker run -d --name nebulo-mongodb -p 27017:27017 -e MONGO_INITDB_DATABASE=nebulo mongo:7.0
	docker run -d --name nebulo-redis -p 6379:6379 redis:7.2-alpine

db-down: ## Stop and remove database containers
	docker stop nebulo-mongodb nebulo-redis || true
	docker rm nebulo-mongodb nebulo-redis || true

# Health checks
health: ## Check server health
	@echo "Checking server health..."
	@curl -f http://localhost:8080/health || echo "Server is not responding"

health-device: ## Check device server health
	@echo "Checking device server health..."
	@curl -f http://localhost:8081/internal/storage || echo "Device server is not responding"

test-api: ## Run comprehensive API tests
	@echo "Running API tests..."
	@chmod +x scripts/test-api.sh
	@./scripts/test-api.sh