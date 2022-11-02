# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BIN_DIR=bin
CMD=exec
CMD_DIR=./cmd

all: module
module:
	@CGO_ENABLED=1 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags "-s -w" -o $(BIN_DIR)/$(CMD) -v $(CMD_DIR)

clean: bin/
	@rm -rf bin/*

