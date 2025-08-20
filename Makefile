MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))



.PHONY: di lint build-api build
all : di lint build-api build
di:
	@echo "Injecting Dependencies"
	@cd $(MAKEFILE_DIR) && go install github.com/google/wire/cmd/wire@latest
	@cd $(MAKEFILE_DIR) && wire ./di

lint:
	@echo "Linting Modules"
	@cd $(MAKEFILE_DIR) && go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
	@cd $(MAKEFILE_DIR) && golangci-lint run

build-api:
	@echo "Building Api executable"
	@cd $(MAKEFILE_DIR) && @go build -o out/api ./cmd/api

build: di lint build-api