NAME    := tumbling
VERSION := v0.1.0
BUILD   := $(shell git rev-parse --short HEAD)
LDFLAGS :=-ldflags "-X main.version=${VERSION} -X main.build=${BUILD}"

.PHONY: build
build:
	go build $(LDFLAGS) -o bin/$(NAME) ./cmd/$(NAME)/main.go

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf vendor/*

.PHONY: setup
setup:
	go get -u github.com/golang/dep/cmd/dep

.PHONY: deps
deps: setup
	dep ensure -update

.PHONY: install
install:
	go install $(LDFLAGS) ./cmd/$(NAME)
