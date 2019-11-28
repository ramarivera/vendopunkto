EXECUTABLE := vendopunkto-server
GITVERSION := $(shell git describe --dirty --always --tags --long)
GOPATH ?= ${HOME}/go
PACKAGENAME := $(shell go list -m -f '{{.Path}}')
MIGRATIONDIR := store/postgres/migrations
MIGRATIONS :=  $(wildcard ${MIGRATIONDIR}/*.sql)
TOOLS := ${GOPATH}/bin/mockery \
    ${GOPATH}/bin/wire

.PHONY: default
default: ${EXECUTABLE}

${GOPATH}/bin/mockery:
	go get github.com/vektra/mockery/internal/cmd/mockery

${GOPATH}/bin/wire:
	go get github.com/google/wire
	go get github.com/google/wire/internal/cmd/wire

tools: ${TOOLS}

internal/cmd/wire_gen.go: internal/cmd/wire.go
	wire ./internal/cmd/...

.PHONY: mocks
mocks: tools
	mockery -dir ./gorestapi -name ThingStore

.PHONY: ${EXECUTABLE}
${EXECUTABLE}: tools internal/cmd/wire_gen.go 
	# Compiling...
	go build -ldflags "-X ${PACKAGENAME}/conf.Executable=${EXECUTABLE} -X ${PACKAGENAME}/conf.GitVersion=${GITVERSION}" -o ${EXECUTABLE}

.PHONY: test
test: 
	go test -cover ./...

.PHONY: deps
deps:
	# Fetching dependancies...
	go get -d -v # Adding -u here will break CI

.PHONY: relocate
relocate:
	@test ${TARGET} || ( echo ">> TARGET is not set. Use: make relocate TARGET=<target>"; exit 1 )
	$(eval ESCAPED_PACKAGENAME := $(shell echo "${PACKAGENAME}" | sed -e 's/[\/&]/\\&/g'))
	$(eval ESCAPED_TARGET := $(shell echo "${TARGET}" | sed -e 's/[\/&]/\\&/g'))
	# Renaming package ${PACKAGENAME} to ${TARGET}
	@grep -rlI '${PACKAGENAME}' * | xargs -i@ sed -i 's/${ESCAPED_PACKAGENAME}/${ESCAPED_TARGET}/g' @
	# Complete... 
	# NOTE: This does not update the git config nor will it update any imports of the root directory of this project.

dbclean:
	docker stop vendopunktopostgres; 
	docker rm vendopunktopostgres;
	docker volume rm vendopunkto_db;
	docker-compose up -d vendopunktopostgres

run:
	STORAGE_HOST=localhost \
	PLUGINS_ENABLED="wallet|http://localhost:3333" \
	${GOPATH}/src/${PACKAGENAME}/vendopunkto-server api

build-monero:
	go build -o ./vendopunkto-monero ./plugins/monero/main.go

run-monero:
	MONERO_WALLET_RPC_URL=http://localhost:18082 \
	${GOPATH}/src/${PACKAGENAME}/vendopunkto-monero