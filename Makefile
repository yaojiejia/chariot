BINARY_NAME=chariot
GOFILES=$(wildcard *.go)
build:
	@echo "Building..."
	go build -o $(BINARY_NAME) $(GOFILES)

run: build
	@echo "Running..."
	./$(BINARY_NAME)

test:
	@echo "Running tests..."
	go test -v ./...	