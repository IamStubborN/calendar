# Go parameters
    BINARY_NAME=calendar

# Main package
    MAIN=./cmd

# Commands
    all:
		go mod tidy
		go test -v ./...
		go build -o $(BINARY_NAME) $(MAIN)
    build:
		go build -o $(BINARY_NAME) -v $(MAIN)
    test:
		go test -v ./...
    clean:
		go clean $(MAIN)
		rm -rf $(BINARY_NAME)
    run:
		go run $(MAIN)
    generate:
		go generate ./...
    lint:
		golangci-lint run ./... -v
    compose_up:
		docker-compose up
    compose_down:
		docker-compose down
