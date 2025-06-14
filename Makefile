.PHONY: build
build:
	mkdir -p bin/gopro-1
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/gopro-1/main ./cmd/gopro-1

.PHONY: clean
clean:
	rm -rf bin/gopro-1/main

.PHONY: lint
lint:
	golangci-lint run