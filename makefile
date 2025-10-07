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
	(cd src/sdlmenu && go build -tags static -v -o ../../$(BIN_DIR)/sdlmenu/sdlmenu)

build-router:
	@echo "building router"
	(cd src/router && go build -v -o ../../$(BIN_DIR)/router/router)

run-sdlmenu:
	@echo "running sdlmenu"
	./build/bin/sdlmenu/sdlmenu

run-router:
	@echo "running router"
	./build/bin/router/router

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BIN_DIR) $(ARTIFACT_DIR)
	@find . -name go.sum -delete

test:
	@for module in $(MODULE_DIRS); do \
		echo "Testing $$module..."; \
		(cd $$module && GO111MODULE=on go test -v ./...); \
	done


vendor:
	@for module in $(MODULE_DIRS); do \
		echo "Vendoring dependencies for $$module..."; \
		(cd $$module && GO111MODULE=on go mod vendor); \
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

