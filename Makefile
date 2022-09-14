# If present, load environment variables from .env
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: image
image:
	docker build . -t devxstats

.PHONY: build
build:
	go build -o bin/devxstats cmd/devxstats/main.go

.PHONY: install
install:
	@echo ">> Installing devxstats locally"
	go install cmd/devxstats/main.go

.PHONY: mongo
mongo:
	docker run --rm -d -p 27017:27017 --name db mongo:latest

.PHONY: run
run:
	go run cmd/devxstats/main.go

.DEFAULT_GOAL := build