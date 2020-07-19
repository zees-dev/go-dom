all: test clean build

clean:
	rm -rf ./go-dom

build:
	go build -o go-dom

test:
	go test -v -short ./...
.PHONY: test
