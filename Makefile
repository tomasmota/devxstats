run:
	pack build devxstats && docker-compose up

run-local:
	ENVIRONMENT='dev' go run main.go

build:
	go build -o bin/devxstats

install:
	@echo ">> Installing devxstats locally"
	go install .
