run:
	ENVIRONMENT='dev' go run main.go

build:
	go build -o bin/devxstats

install:
	@echo ">> Installing devxstats locally"
	go install .
