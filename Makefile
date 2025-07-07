.PHONY: build test lint clean run

# Build variables
BINARY_NAME=task-manager
MAIN_FILE=cmd/task-manager/main.go

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
	$(GOGET) github.com/rs/zerolog
	$(GOGET) github.com/fatih/color
	$(GOGET) github.com/stretchr/testify

.DEFAULT_GOAL := build