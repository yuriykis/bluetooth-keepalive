BINARY_NAME=bth-speaker-on

build: export GOOS=darwin
build: export GOARCH=amd64
build:
	@go build -o bin/$(BINARY_NAME) -v

run: build
	@./bin/$(BINARY_NAME)