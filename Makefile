BINARY_NAME=bth-speaker-on
UP_INTERVAL=5

build: export GOOS=darwin
build: export GOARCH=arm64
build:
	@go build -o bin/$(BINARY_NAME) -v

run: build
	@./bin/$(BINARY_NAME) start up-interval=$(UP_INTERVAL)

run-default: build
	@./bin/$(BINARY_NAME) start
	
install: build
	@go install