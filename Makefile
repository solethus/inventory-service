.PHONY: all build test clean proto

all: build

build:
	go build -o bin/inventory-service ./cmd/inventory

test:
	go test ./...

clean:
	rm -rf bin

proto:
	./scripts/generate_proto.sh

run:
	go run ./main.go
