# Build directories
ROOT := ./
MODULES := $(ROOT)/src/sdlmenu/ $(ROOT)/src/router/ 
					#	$(ROOT) \
					#	$(ROOT) 
						
BIN_DIR := build/bin
ARTIFACTS_DIR := build/artifacts

GO_PROXY := direct
TAG ?= latest

SDLMENU := github.com/epxzzy/k-7/sdlmenu
ROUTER := github.com/epxzzy/k-7/router
TEST_DIR := src/test


.PHONY: build clean test run
	
build: build-sdlmenu build-router

build-sdlmenu:
	@echo "building sdlmenu"
	(cd src/sdlmenu && go build -tags static -o ../../$(BIN_DIR)/sdlmenu/sdlmenu.exe)

build-router:
	@echo "building router"
	(cd src/router && go build -o ../../$(BIN_DIR)/router/router.exe)


clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BIN_DIR) $(ARTIFACT_DIR)
	@find . -name go.sum -delete

test:
	@for module in $(MODULE_DIRS); do \
		echo "Testing $$module..."; \
		(cd $$module && GO111MODULE=on go test -v ./...); \
	done

run-router: build
	@echo "Starting router module..."
	@$(BIN_DIR)/router

vendor:
	@for module in $(MODULE_DIRS); do \
		echo "Vendoring dependencies for $$module..."; \
		(cd $$module && GO111MODULE=on go mod vendor); \
	done

# Cross-compilation 
build-linux:
	@for module in $(MODULE_DIRS); do \
		echo "Building $$module for Linux..."; \
		(cd $$module && \
			GOOS=linux GOARCH=amd64 GO111MODULE=on go build \
			-o $(BIN_DIR)/linux-amd64/$$(basename $$module) \
			.); \
	done

# Dependency resolution between modules
$(BIN_DIR)/router: $(BIN_DIR)/sdlmenu $(BIN_DIR)/cli

# Module installation (if needed as dependencies)
install-modules:
	@echo "Installing modules to GOPATH..."
	@for module in $(MODULE_DIRS); do \
		echo "Installing $$module..."; \
		(cd $$module && GO111MODULE=on go install); \
	done

