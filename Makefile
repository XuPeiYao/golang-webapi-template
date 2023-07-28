init:
	@$(MAKE) setup
	@$(MAKE) restore

setup:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/swaggo/swag/cmd/swag@latest

restore:
	go mod tidy

vet:
	@$(MAKE) vet.api-server

vet.api-server:
	go vet ./cmd/api-server

build:
	@$(MAKE) build.api-server

build.api-server:
	bash ./scripts/wire.sh
	bash ./scripts/update-swagger.sh
	bash ./scripts/build.sh
