# Build directories
BUILD_DIR := build
BIN_DIR := $(BUILD_DIR)/bin
ARTIFACTS_DIR := $(BUILD_DIR)/artifacts

# Source directories
SRC_DIR := src/main/
TEST_DIR := src/test/

# Application
APP_NAME := main.exe

.PHONY: build clean test run

build:
	@echo "Building application..."
	@cd $(SRC_DIR) && \
		go build -tags static -o ../../$(BIN_DIR)/$(APP_NAME) ./

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)

test:
	@echo "Running tests..."
	@cd $(SRC_DIR) && \
		go test -v ./...

run: build
	@echo "Starting application..."
	@./$(BIN_DIR)/$(APP_NAME)

vendor:
	@cd $(SRC_DIR) && \
		go mod tidy && \
		go mod vendor

# Cross-compilation example
build-linux:
	@cd $(SRC_DIR) && \
		GOOS=linux GOARCH=amd64 go build -o ../../$(BIN_DIR)/$(APP_NAME)-linux ./cmd/app
