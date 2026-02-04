.PHONY: build run
build:
# 	bin/hexlet-path-size
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

run:
	go run ./cmd/hexlet-path-size/

test:
	go test -v ./internal/pathsize/

lint:
	golangci-lint run