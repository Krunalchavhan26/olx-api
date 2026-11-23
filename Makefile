# Define variables
CMD_DIR = ./cmd

.PHONY: bin run

build:
	@go build -o bin/api ${CMD_DIR}/api/

run: build
	@./bin/api