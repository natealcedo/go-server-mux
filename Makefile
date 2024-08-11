.PHONY: build clean create fix install lint up run start test
APP_NAME = app
BIN = bin/$(APP_NAME)

run:
	@go run main.go

start: build
	@$(BIN)

install:
	@go mod tidy

build:
	@go build -o $(BIN)

clean:
	@rm -rf bin/*

test:
	@go test -v ./...  # -count=1 to avoid caching

lint:
	@gofmt -l .

fix:
	@gofmt -w .

create:
	@name=$${name}; \
	GOOSE_MIGRATION_DIR="./migrations" goose postgres user=postgres dbname=postgres create $$name sql

up:
	@ GOOSE_MIGRATION_DIR="./migrations" goose postgres "user=postgres dbname=postgres" up

down:
	@ GOOSE_MIGRATION_DIR="./migrations" goose postgres "user=postgres dbname=postgres" down
