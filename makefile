# Go parameters
BINARY_NAME=balsim

all: test build
build: 
	go build -o $(BINARY_NAME) -v
test: 
	go test -count=1 -v ./...
clean: 
	go clean
	rm -f $(BINARY_NAME)
