image:
	docker build . -t devxstats

build:
	go build -o bin/devxstats cmd/devxstats/main.go

install:
	@echo ">> Installing devxstats locally"
	go install cmd/devxstats/main.go
