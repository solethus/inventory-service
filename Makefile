.PHONY: all build test clean proto

all: build

build:
	go build -o bin/inventory-service .

test:
	go test ./...

clean:
	rm -rf bin

run:
	go run ./main.go
