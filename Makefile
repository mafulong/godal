PROJECT_NAME := "github.com/mafulong/godal"
PKG := "$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v test)
#GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go | grep -v "$(PROJECT_NAME)"/utils | grep -v "$(PROJECT_NAME)"/test)

.PHONY: all dep lint vet test test-coverage build clean

all: build

dep: ## Get the dependencies
	@go mod download

lint: ## Lint Golang files
	@golint -set_exit_status ${PKG_LIST}

vet: ## Run go vet
	@go vet ${PKG_LIST}

test: ## Run unittests
	@go test -short ${PKG_LIST}

test-coverage: ## Run tests with coverage
	@echo ${PKG_LIST}
	@go test -cover -coverprofile=coverage.txt -coverpkg="./..." ./...

build: dep ## Build the binary file
	@go build -i -o build/main $(PKG)

clean: ## Remove previous build
	@rm -f ./build

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'