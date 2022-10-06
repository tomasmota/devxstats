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
	docker run -p 5432:5432 --name pg -e POSTGRES_PASSWORD=${PGPASSWORD} -e POSTGRES_USER=${PGUSER} -e POSTGRES_DB=${PGDATABASE} -d postgres
	sleep 2
	cat migrations/init.sql | docker exec -i pg psql -U ${PGUSER} -d ${PGDATABASE}

.PHONY: run
run:
	go run cmd/devxstats/main.go --octopus

.DEFAULT_GOAL := build