.DEFAULT_GOAL := test

# Clean up
clean:
	@rm -fR ./cover.*
.PHONY: clean

# Run tests and generates html coverage file
cover: test
	@go tool cover -html=./coverage.text -o ./coverage.html
.PHONY: cover

# Download dependencies
depend:
	go mod tidy
.PHONY: depend

# Format all go files
fmt:
	gofmt -s -w -l $(shell go list -f {{.Dir}} ./...)
.PHONY: fmt

# Run tests
test:
	@go test -v -race -coverprofile=./coverage.text -covermode=atomic $(shell go list ./...)
.PHONY: test
