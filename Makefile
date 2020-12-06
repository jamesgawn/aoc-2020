OS=$(shell go env GOOS)
ARCH=$(shell go env GOARCH)
FILENAME=./bin/aoc2020-$(OS)-$(ARCH)

build:
	go build -o ./bin/aoc2020 ./

buildWithArch:
	go build -o $(FILENAME) ./

buildName:
	@echo $(FILENAME)

test:
	go test ./...