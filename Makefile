PHONY: run build

run:
	go run ./cmd
build:
	go build -o bin/lnk-api ./cmd
