BINARY_NAME=selector
 
build:
	go mod tidy
	go build -o build/${BINARY_NAME} main.go
 
clean:
	go clean
	rm build/${BINARY_NAME}