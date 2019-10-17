# メタ情報
NAME := gosimplehttpd
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' \
	   -X 'main.revision=$(REVISION)'
OS := $(shell uname)

# 必要なツール類をセットアップする
## Setup
setup:
	go get github.com/Masterminds/glide
	go get golang.org/x/lint/golint
	go get golang.org/x/tools/cmd/goimports
	go get github.com/Songmu/make2help/cmd/make2help

# glideを使って依存パッケージをインストールする
## Install dependencies
deps: setup
	glide install

## Update dependencies
update: setup
	glide update

## Lint
lint: setup
	go vet $$(glide novendeor)
	for pkg in $$(glide novendor -x); do \
		golint -set_exit_status $$pkg || exit $$?; \
	done

## Format source codes
fmt: setup
	goimports -w $$(glide nv -x)

## build binaries ex. make bin/gosimplehttpd
bin/%: cmd/%/main.go deps
	go build -ldflags "$(LDFLAGS)" -o $@ $<
	mkdir -p bin/$(OS)
	mv $@ bin/$(OS)/

## Show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: setup deps update test lint help

