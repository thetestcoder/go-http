.PHONY: build test lint clean run

# Build variables
BINARY_NAME=basic-http
MAIN_FILE=cmd/basic-http/main.go

# Go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Linting
GOLINT=golangci-lint

all: lint test build

build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_FILE)

test:
	$(GOTEST) -v ./tests/...

lint:
	$(GOLINT) run

clean:
	rm -f $(BINARY_NAME)
	rm -f *.log

run: build
	./$(BINARY_NAME)

deps:
	$(GOMOD) download

.DEFAULT_GOAL := build