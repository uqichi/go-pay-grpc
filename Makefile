.DEFAULT_GOAL := help

protoc: ## Exec protoc
	protoc --go_out=plugins=grpc:. proto/*.proto

server: ## Run server
	go run server.go

test: ## Exec test
	go test -v .

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

