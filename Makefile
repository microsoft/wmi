# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.
GOCMD=go
GOBUILD=$(GOCMD) build -v 
GOHOSTOS=$(strip $(shell $(GOCMD) env get GOHOSTOS))

TAG ?= $(shell git describe --tags)
COMMIT ?= $(shell git describe --always)
BUILD_DATE ?= $(shell date -u +%m/%d/%Y)

PKG := 

all: vendor format library
 
.PHONY: vendor
vendor:
	GO111MODULE=on go mod tidy

library: vendor
	GO111MODULE=on GOARCH=amd64 GOOS=windows $(GOBUILD) ./...

format: vendor
	gofmt -s -w pkg

test: vendor
	GOOS=windows GO111MODULE=on  GOARCH=amd64 go test -v ./go/...
