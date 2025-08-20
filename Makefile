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

.PHONY: clean di lint build-api build dev-nt
all : di lint build-api build

clean:
	@echo Cleaning directories...
	@cd $(MAKEFILE_DIR) && $(RM) $(OUTPUT_DIR)
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
	@cd $(MAKEFILE_DIR) && go build -o out/api$(EXT) ./cmd/api

build: di lint build-api

dev-nt: clean build-api
	@cd $(MAKEFILE_DIR) && .\out\api$(EXT) 

dev: clean build-api
	@cd $(MAKEFILE_DIR) && ./out/api