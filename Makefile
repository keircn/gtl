BINARY_NAME := gtl
PACKAGE := github.com/keircn/gtl
MAIN_PACKAGE := ./cmd/gtl
BUILD_DIR := ./build
DIST_DIR := ./dist

VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GO_VERSION := $(shell go version | cut -d' ' -f3)

LDFLAGS := -ldflags "\
	-s -w \
	-X '$(PACKAGE)/pkg/version.Version=$(VERSION)' \
	-X '$(PACKAGE)/pkg/version.GitCommit=$(GIT_COMMIT)' \
	-X '$(PACKAGE)/pkg/version.BuildDate=$(BUILD_DATE)' \
	-X '$(PACKAGE)/pkg/version.GoVersion=$(GO_VERSION)'"

BUILD_FLAGS := -trimpath
DEV_FLAGS := -race

PLATFORMS := \
	linux/amd64 \
	linux/arm64 \
	darwin/amd64 \
	darwin/arm64 \
	windows/amd64 \
	windows/arm64

.PHONY: all
all: clean format lint test build

.PHONY: deps
deps:
	go mod download
	go mod verify

.PHONY: format
format:
	gofmt -s -w .
	goimports -w . 2>/dev/null || true

.PHONY: lint
lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		go vet ./...; \
		go fmt ./...; \
	fi

.PHONY: test
test:
	go test -v ./...

.PHONY: build
build: deps
	@mkdir -p $(BUILD_DIR)
	go build $(BUILD_FLAGS) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PACKAGE)
	@echo "Built: $(BUILD_DIR)/$(BINARY_NAME)"

.PHONY: install
install: build
	go install $(BUILD_FLAGS) $(LDFLAGS) $(MAIN_PACKAGE)

.PHONY: cross-compile
cross-compile: deps
	@mkdir -p $(BUILD_DIR)
	@for platform in $(PLATFORMS); do \
		OS=$$(echo $$platform | cut -d'/' -f1); \
		ARCH=$$(echo $$platform | cut -d'/' -f2); \
		BINARY_SUFFIX=""; \
		if [ "$$OS" = "windows" ]; then \
			BINARY_SUFFIX=".exe"; \
		fi; \
		OUTPUT="$(BUILD_DIR)/$(BINARY_NAME)-$$OS-$$ARCH$$BINARY_SUFFIX"; \
		echo "  Building for $$OS/$$ARCH..."; \
		GOOS=$$OS GOARCH=$$ARCH go build $(BUILD_FLAGS) $(LDFLAGS) -o $$OUTPUT $(MAIN_PACKAGE); \
		if [ $$? -eq 0 ]; then \
			echo "    $$OUTPUT"; \
		else \
			echo "    Failed to build for $$OS/$$ARCH"; \
		fi; \
	done

.PHONY: package
package: cross-compile
	@mkdir -p $(DIST_DIR)
	@for platform in $(PLATFORMS); do \
		OS=$$(echo $$platform | cut -d'/' -f1); \
		ARCH=$$(echo $$platform | cut -d'/' -f2); \
		BINARY_SUFFIX=""; \
		ARCHIVE_EXT="tar.gz"; \
		if [ "$$OS" = "windows" ]; then \
			BINARY_SUFFIX=".exe"; \
			ARCHIVE_EXT="zip"; \
		fi; \
		BINARY_PATH="$(BUILD_DIR)/$(BINARY_NAME)-$$OS-$$ARCH$$BINARY_SUFFIX"; \
		if [ -f "$$BINARY_PATH" ]; then \
			PACKAGE_NAME="$(BINARY_NAME)-$(VERSION)-$$OS-$$ARCH"; \
			PACKAGE_DIR="$(DIST_DIR)/$$PACKAGE_NAME"; \
			mkdir -p "$$PACKAGE_DIR"; \
			cp "$$BINARY_PATH" "$$PACKAGE_DIR/$(BINARY_NAME)$$BINARY_SUFFIX"; \
			cp README.md "$$PACKAGE_DIR/" 2>/dev/null || true; \
			cp LICENSE "$$PACKAGE_DIR/" 2>/dev/null || true; \
			if [ "$$ARCHIVE_EXT" = "zip" ]; then \
				cd $(DIST_DIR) && zip -r "$$PACKAGE_NAME.zip" "$$PACKAGE_NAME/" > /dev/null; \
				echo "  $$PACKAGE_NAME.zip"; \
			else \
				cd $(DIST_DIR) && tar -czf "$$PACKAGE_NAME.tar.gz" "$$PACKAGE_NAME/" 2>/dev/null; \
				echo "  $$PACKAGE_NAME.tar.gz"; \
			fi; \
			rm -rf "$$PACKAGE_DIR"; \
		fi; \
	done

.PHONY: release
release: clean all package
	@echo "Version: $(VERSION)"
	@echo "Packages created in: $(DIST_DIR)"
	@ls -la $(DIST_DIR)/

.PHONY: run
run: build
	$(BUILD_DIR)/$(BINARY_NAME) $(ARGS)

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR) $(DIST_DIR)
	rm -f coverage.out coverage.html
	go clean -cache -testcache -modcache 2>/dev/null || true

.PHONY: version
version:
	@echo "Version: $(VERSION)"
	@echo "Git Commit: $(GIT_COMMIT)"
	@echo "Build Date: $(BUILD_DATE)"
	@echo "Go Version: $(GO_VERSION)"

.PHONY: watch
watch: ## requires entr
	@if command -v entr >/dev/null 2>&1; then \
		echo "Watching for changes..."; \
		find . -name "*.go" | entr -r make run-dev ARGS="$(ARGS)"; \
	else \
		echo "entr not found."; \
	fi

.DEFAULT_GOAL := build

