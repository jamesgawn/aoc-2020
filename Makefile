OS=$(shell go env GOOS)
ARCH=$(shell go env GOARCH)
FILENAME=./bin/$(OS)-$(ARCH)/aoc2020

build:
	go build -o ./bin/aoc2020 ./

buildWithArch:
	go build -o $(FILENAME) ./

test:
	go test ./...