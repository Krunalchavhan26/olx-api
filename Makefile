# Define variables
CMD_DIR = ./cmd

.PHONY: bin run

build:
	@go build -o bin/api ${CMD_DIR}/api/

run: build
	@./bin/api

migrate-up:
	@go run ${CMD_DIR}/migrate up

migrate-down:
	@go run ${CMD_DIR}/migrate down