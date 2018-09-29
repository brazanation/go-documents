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
	@go get -u gopkg.in/alecthomas/gometalinter.v2
	@gometalinter.v2 --install
.PHONY: depend

# Format all go files
fmt:
	gofmt -s -w -l $(shell go list -f {{.Dir}} ./...)
.PHONY: fmt

# Run linters
lint:
	@gometalinter.v2 \
		--disable-all \
		--exclude=vendor \
		--deadline=180s \
		--enable=gofmt \
		--linter='errch:errcheck {path}:PATH:LINE:MESSAGE' \
		--enable=errch \
		--enable=vet \
		--enable=gocyclo \
		--cyclo-over=15 \
		--enable=golint \
		--min-confidence=0.85 \
		--enable=ineffassign \
		--enable=misspell \
		./..
.PHONY: lint

# Run tests
test:
	@go test -v -race -coverprofile=./coverage.text -covermode=atomic $(shell go list ./...)
.PHONY: test
