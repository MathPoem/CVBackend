APP_BUILD = build/apiserver

.DEFAULT_GOAL = build

.PHONY: build
build:
	clear; go build -o ${APP_BUILD} -v ./cmd/main.go

.PHONY: start
start: build
	${APP_BUILD}

.PHONY: migrate-dev
migrate-dev:
	migrate -path schema -database "postgres://postgres:fdaf2324h908n0lkda@localhost:5436/postgres?sslmode=disable" ${type}
