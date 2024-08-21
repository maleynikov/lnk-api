PHONY: run bin

run:
	go run ./cmd
bin:
	go build -o bin/lnk-api ./cmd
