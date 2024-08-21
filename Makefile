.PHONY: run bin

run:
	go run ./cmd
bin:
	rm -rf ./bin && go build -o bin/lnk-api ./cmd
