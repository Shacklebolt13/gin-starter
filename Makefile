MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
OUTPUT_DIR := .\\out

ifeq ($(OS),Windows_NT)
    RM = rmdir /S /Q
    MKDIR = mkdir
	EXT = ".exe"
else
    RM = rm -rf
    MKDIR = mkdir -p
	EXT = ""
endif

.PHONY: tidy clean di lint swag build-api build dev-nt
all : tidy di lint build-api build swag

clean:
	@echo Cleaning directories...
	@cd $(MAKEFILE_DIR) && $(RM) $(OUTPUT_DIR)
di:
	@echo Injecting Dependencies
	@go install github.com/google/wire/cmd/wire@latest
	@cd $(MAKEFILE_DIR) && wire gen ./di

swag:
	@echo Generating Swagger Docs
	@go install github.com/swaggo/swag/cmd/swag@latest
	@cd $(MAKEFILE_DIR) && swag init --dir cmd/api

tidy:
	@cd $(MAKEFILE_DIR) && go mod tidy

lint:
	@echo Linting Modules
	@go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
	@cd $(MAKEFILE_DIR) && golangci-lint-v2 run ./...

build-api:
	@echo Building Api executable
	@cd $(MAKEFILE_DIR) && go build -o out/api$(EXT) ./cmd/api

build: di swag build-api

dev: clean build-api
ifeq ($(OS),Windows_NT)
	@cd $(MAKEFILE_DIR) && .\out\api$(EXT) 
else
	@cd $(MAKEFILE_DIR) && ./out/api	
endif
