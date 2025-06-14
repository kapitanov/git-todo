GOIMPORTS     := v0.33.0
GOTESTSUM     := v1.12.2
GOLANGCI_LINT := v2.1.6

GO_FILES      := $(shell find . -type f -name '*.go' -not -path './artifacts/*' -not -path "./vendor/*")
LOCAL_PACKAGE := $(shell go list -m | head -n 1)

.PHONY: all build install uninstall format test lint release website
all: build test lint

build:
	@mkdir -p artifacts
	go build -o artifacts/git-todo .

install:
	go install -v

uninstall:
	@LOCAL_BINARY="$$(go env GOPATH)/bin/git-todo"; \
	if [ -f "$$LOCAL_BINARY" ]; then \
		rm -f "$$LOCAL_BINARY"; \
		echo "Uninstalled $$LOCAL_BINARY"; \
	else \
		echo "No binary found at $$LOCAL_BINARY"; \
	fi

format:
	go run golang.org/x/tools/cmd/goimports@$(GOIMPORTS) --local $(LOCAL_PACKAGE) -w -format-only $(GO_FILES)

test:
	@mkdir -p artifacts
	go run gotest.tools/gotestsum@$(GOTESTSUM) --junitfile=artifacts/junit.xml -- -coverprofile=artifacts/coverage.out -covermode=count ./...
	go tool cover -html=artifacts/coverage.out -o artifacts/coverage.html
	@go tool cover -func=./artifacts/coverage.out | grep 'total' | awk '{printf "%s\n", $$3}'

lint:
	@mkdir -p artifacts
	go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT) run --output.code-climate.path ./artifacts/lint.json --output.tab.path stdout

release:
	goreleaser release --snapshot --clean

website:
	docker run --rm -t -v "$${PWD}:/mnt" -w /mnt/website squidfunk/mkdocs-material:latest -- build
