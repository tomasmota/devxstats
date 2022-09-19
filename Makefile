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

.PHONY: pg
pg:
	docker run --name pg -e POSTGRES_PASSWORD=mypass -e POSTGRES_USER=myuser -e POSTGRES_DB=devxstats -d postgres
#TODO: Put in .sql file here 
.PHONY: run
run:
	go run cmd/devxstats/main.go

.DEFAULT_GOAL := build