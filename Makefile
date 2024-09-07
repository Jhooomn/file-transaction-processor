APP_NAME := file-transaction-processor
BUILD_DIR := ./build

GO := go
GO_BUILD := $(GO) build
GO_TEST := $(GO) test
GO_RUN := $(GO) run
GO_TIDY := $(GO) mod tidy
GO_VENDOR := $(GO) mod vendor
GO_VET := $(GO) vet ./...
GO_FMT := $(GO) fmt ./...


build:
	@echo "Building $(APP_NAME)..."
	$(GO_BUILD) -o $(BUILD_DIR)/$(APP_NAME) .

run:
	@echo "Running $(APP_NAME)..."
	$(GO_RUN) .

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
	$(GO_VET)

mocks:
	mockery --all --recursive --output=mocks

local_pipeline:
	make tidy_vendor && make format && make build && make test && make mocks && make run 

.PHONY: build run test tidy_vendor format local_pipeline  mocks
