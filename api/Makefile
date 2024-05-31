LOCAL_BIN:=$(CURDIR)/bin

GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
GOLANGCI_TAG:=1.56.2
MOCKGEN_TAG:=latest
SMART_IMPORTS := ${LOCAL_BIN}/smartimports


.PHONY: install-swag
install-swag:
	go install github.com/swaggo/swag/cmd/swag@latest


.PHONY: create-swag
create-swag:
	swag init -g ./cmd/main.go  -o ./internal/docs

# build app
.PHONY: build
build:
	go mod download && CGO_ENABLED=0  go build \
		-o ./bin/main$(shell go env GOEXE) ./cmd/main.go

# precommit jobs
.PHONY: precommit
precommit: format lint


# run diff lint like in pipeline
.PHONY: .lint
.lint: install-lint
	$(info Running lint...)
	$(GOLANGCI_BIN) run --new-from-rev=origin/master --config=.golangci.yaml ./...

# golangci-lint diff master
.PHONY: lint
lint: .lint

# run full lint like in pipeline
.PHONY: .lint-full
.lint-full: install-lint
	$(GOLANGCI_BIN) run --config=.golangci.yaml ./...

# golangci-lint full
.PHONY: lint-full
lint-full: .lint-full


ifeq (,$(wildcard .env))
    # Если файл .env отсутствует, устанавливаем параметры по умолчанию
    POSTGRES_USER := 'SuperPuperUser'
    POSTGRES_PASSWORD := 'SuperStrongPassword'
    POSTGRES_DB := 'dev'
    POSTGRES_HOST := '192.168.1.70'
    POSTGRES_PORT := 5432
else
    # Иначе, подключаем переменные из файла .env
    include .env
    export
endif
POSTGRES_SETUP_TEST := user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} host=${POSTGRES_HOST} port=${POSTGRES_PORT} sslmode=disable

INTERNAL_PKG_PATH=$(CURDIR)/internal
MIGRATION_FOLDER=$(CURDIR)/migrations

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: migration-up
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up


.PHONY: migration-down
migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down

.PHONY: docker-build
docker-build:
	docker build . -t ghcr.io/egortarasov/lct-2024/api

.PHONY: docker-push
docker-push:
	docker push ghcr.io/egortarasov/lct-2024/api:latest