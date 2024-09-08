APP_NAME := file-transaction-processor
BUILD_DIR := ./build

GO := go
GO_BUILD := $(GO) build
GO_TEST := $(GO) test
GO_TIDY := $(GO) mod tidy
GO_VENDOR := $(GO) mod vendor
GO_FMT := $(GO) fmt ./...

build:
	mkdir -p dist
	env GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -tags lambda.norpc -ldflags="-s -w" -o ./dist/bootstrap main.go
	chmod +x ./dist/bootstrap
	cd ./dist && zip -FS bootstrap.zip bootstrap ../.env ../data -r && cd ..

test:
	@echo "Running tests..."
	go test -count=1 ./...

tidy_vendor:
	rm -rf vendor
	@echo "Tidying and vendoring dependencies..."
	$(GO_TIDY)
	$(GO_VENDOR)

format:
	@echo "Formatting current golang code..."
	$(GO_FMT)

## WNS
mock:
	mockery --all --dir=./processor --output=mocks 
	mockery --all --dir=./infrastructure/email --output=mocks 
	mockery --all --dir=./infrastructure/repository --output=mocks 

local_pipeline:
	make tidy_vendor && make format && make test 

.PHONY: build test tidy_vendor format mock local_pipeline 
