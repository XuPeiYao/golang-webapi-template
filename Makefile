.DEFAULT_GOAL := ci.build

ci.build: init build

init: setup restore

setup: ## Setup required CLI tools
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/swaggo/swag/cmd/swag@latest

restore: ## Restore reqired packages
	go mod tidy

vet: vet.api-server

vet.api-server:
	go vet ./cmd/api-server

build:  build.api-server ## Build api-server

build.api-server: 
	bash ./scripts/api-server.wire.sh
	bash ./scripts/api-server.swag.sh
	bash ./scripts/api-server.build.sh

help: ## Show help message.
	@printf "Usage:\n"
	@printf "  make <target>\n\n"
	@printf "Targets:\n"
	@perl -nle'print $& if m{^[a-zA-Z0-9_-]+:.*?## .*$$}' $(MAKEFILE_LIST) | \
		sort | \
		awk 'BEGIN {FS = ":.*?## "}; \
		{printf "  %-18s %s\n", $$1, $$2}'