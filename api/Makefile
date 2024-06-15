LOCAL_BIN:=$(CURDIR)/bin

GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
GOLANGCI_TAG:=1.59.0
MOCKGEN_TAG:=latest
SMART_IMPORTS := ${LOCAL_BIN}/smartimports


# install golangci-lint binary
.PHONY: install-lint
install-lint:
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info Downloading golangci-lint v$(GOLANGCI_TAG))
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_TAG)
endif

.PHONY: format
format:
	$(info Running goimports...)
	test -f ${SMART_IMPORTS} || GOBIN=${LOCAL_BIN} go install github.com/pav5000/smartimports/cmd/smartimports@latest
	${SMART_IMPORTS} -exclude pkg/,internal/pb  -local 'github.com'


# precommit jobs
.PHONY: precommit
precommit: format lint


# run diff lint like in pipeline
.PHONY: .lint
.lint: install-lint
	$(info Running lint...)
	$(GOLANGCI_BIN) run --new-from-rev=origin/main --config=.golangci.yaml ./...

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


.PHONY:go-create-grpc
go-create-grpc:
	protoc \
	--go_out=./internal/stubs/ \
	--go-grpc_out=./internal/stubs/ \
	--proto_path ../proto/  inference.proto


ifeq (,$(wildcard .env))
    # Если файл .env отсутствует, устанавливаем параметры по умолчанию
    POSTGRES_USER := 'SuperPuperUser'
    POSTGRES_PASSWORD := 'SuperStrongPassword'
    POSTGRES_DB := 'dev'
    POSTGRES_HOST := '127.0.0.1'
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



VERSION := 0.0.1


.PHONY: docker-build-amd64
docker-build-amd64:
	docker build \
		--tag ghcr.io/egortarasov/lct-2024/api:amd64-$(VERSION) \
		--platform linux/amd64 .

.PHONY: docker-build-arm64
docker-build-arm64:
	docker build \
		--tag ghcr.io/egortarasov/lct-2024/api:arm64-$(VERSION) \
		--platform linux/arm64 .



.PHONY: docker-push-amd64
docker-push-amd64:
	docker push ghcr.io/egortarasov/lct-2024/api:amd64-$(VERSION)

.PHONY: docker-push-arm64
docker-push-arm64:
	docker push ghcr.io/egortarasov/lct-2024/api:arm64-$(VERSION)

.PHONY: docker-publish-amd64
docker-publish-amd64: docker-build-amd64 docker-push-amd64

.PHONY: docker-publish-arm64
docker-publish-arm64: docker-build-arm64 docker-push-arm64
