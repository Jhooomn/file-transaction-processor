APP_NAME := file-transaction-processor
BUILD_DIR := ./build

GO := go
GO_BUILD := $(GO) build
GO_TEST := $(GO) test
GO_TIDY := $(GO) mod tidy
GO_VENDOR := $(GO) mod vendor
GO_FMT := $(GO) fmt ./...

build:
	@echo "Building $(APP_NAME)..."
	$(GO_BUILD) -o $(BUILD_DIR)/$(APP_NAME) .


test:
	@echo "Running tests..."
	$(GO_TEST) ./...


tidy_vendor:
	@echo "Tidying and vendoring dependencies..."
	$(GO_TIDY)
	$(GO_VENDOR)

format:
	@echo "Formatting current golang code..."
	$(GO_FMT)

mock:
	mockery --all --dir=./processor --output=mocks 

local_pipeline:
	make tidy_vendor && make mock && make format && make build && make test 

.PHONY: build test tidy_vendor format mock local_pipeline 
