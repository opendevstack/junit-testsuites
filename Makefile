SHELL = /bin/bash
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

##@ General

# help target is based on https://github.com/operator-framework/operator-sdk/blob/master/release/Makefile.
.DEFAULT_GOAL := help
help: ## Show this help screen.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help

##@ Building

build: ## Build Go binary.
	cd cmd/combine-junit-testsuites && CGO_ENABLED=0 go build
.PHONY: build

##@ Testing

test: test-cmd test-pkg ## Run complete testsuite.
.PHONY: test

test-cmd: ## Run testsuite of commands.
	go test -v ./cmd/...
.PHONY: test-cmd

test-pkg: ## Run testsuite of packages.
	go test -v -v ./combine/...
.PHONY: test-pkg
