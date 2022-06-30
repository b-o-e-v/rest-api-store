.PHONY: build
build:
	go build -v ./cmd/server

.PHONY: run
run:
	go run ./cmd/server

.DEFAULT_GOAL := run
