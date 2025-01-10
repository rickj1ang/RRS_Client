# Makefile for the CLI App

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary name
BINARY_NAME=rrsc

# Main package path
MAIN_PACKAGE=.

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PACKAGE)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run: build
	./$(BINARY_NAME)

deps:
	$(GOGET) -v -t -d ./...

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)_linux -v $(MAIN_PACKAGE)

build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME).exe -v $(MAIN_PACKAGE)

build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME)_mac -v $(MAIN_PACKAGE)

# Docker
docker-build:
	docker build -t $(BINARY_NAME) .

.PHONY: all build test clean run deps build-linux build-windows build-mac docker-build

