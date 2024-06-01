BIN := "./bin/bfa-protection-server"
GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build: clean ## Builds application
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/bfa-protection

clean:
	rm -rf $(BIN)

compile-proto: ## Generates proto code
	protoc --go_out=pkg/server/grpc/ --go_opt=paths=source_relative \
        --go-grpc_out=pkg/server/grpc/ --go-grpc_opt=paths=source_relative \
        api/grpc/bfa_protection.proto

help: pad = 24 # padding for two columns
help:	## Show this help
	@echo
	@echo "The Project."
	@echo
	@echo "Commands:"
	@fgrep -h "##" $(MAKEFILE_LIST) \
		| fgrep -v fgrep \
		| sed -e 's/^/  /' -e 's/:/ /' -e 's/	//g' \
		| sort -k 1 \
		| grep -v '^  #' \
		| awk -F "#" '{printf ("%s% *s%s\n", $$1, $(pad)-length($$1), "", $$3)}'
	@echo

lint:
	go install golang.org/x/lint/golint@latest
	golint ./...

migrate-down: ## Run migrations down
	migrate -database postgres://postgres@localhost:5432/bfa_protection?sslmode=disable -path ./db/migrations down

migrate-up: ## Run migrations upfront
	migrate -database postgres://postgres@localhost:5432/bfa_protection?sslmode=disable -path ./db/migrations up

run: build ## Run application
	$(BIN)

test: ## Run tests
	$(BIN) > /dev/null &
	go test ./... ${GO_TEST_OPTIONS}

test-race:	## Run tests with race checks
	$(BIN) > /dev/null &
	go test -race ./... ${GO_TEST_OPTIONS}

.PHONY: build clean compile-proto help lint migrate-down migrate-up run test test-race