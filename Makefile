# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.
GOCMD=go
GOBUILD=$(GOCMD) build -v 
GOHOSTOS=$(strip $(shell $(GOCMD) env get GOHOSTOS))

TAG ?= $(shell git describe --tags)
COMMIT ?= $(shell git describe --always)
BUILD_DATE ?= $(shell date -u +%m/%d/%Y)

PKG := 

all: tidy format library
 
.PHONY: vendor
.PHONY: tidy
tidy:
	-go mod tidy -e

vendor:
	GO111MODULE=on go mod tidy

library:
	GO111MODULE=on GOARCH=amd64 GOOS=windows $(GOBUILD) ./...

format:
	gofmt -s -w pkg

test:
	GOOS=windows GO111MODULE=on  GOARCH=amd64 go test -v ./go/...
