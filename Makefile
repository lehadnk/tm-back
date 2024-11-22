.PHONY: test

test:
	go test ./test/...

build:
	go mod tidy
	go build -o tmback ./src