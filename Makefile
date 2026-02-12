.PHONY: build run test lint install
build:
# 	bin/hexlet-path-size
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

install:
	go install ./cmd/hexlet-path-size

run:
	go run ./cmd/hexlet-path-size/

test:
	go test -v .
lint:
	golangci-lint run