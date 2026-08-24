GO ?= go

.PHONY: test race vet build run seed-user measure docker-build

test:
	$(GO) test ./... -count=1

race:
	$(GO) test -race ./... -count=1

vet:
	$(GO) vet ./...

build:
	$(GO) build ./...

run:
	$(GO) run ./cmd/server

seed-user:
	$(GO) run ./cmd/seed-user

measure:
	$(GO) run ../../.agents/skills/go-base-project-create/scripts/measure_project.go -root . -enforce

docker-build:
	docker build -t agent-runtime-orchestrator:local .
