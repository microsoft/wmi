# Copyright (c) Microsoft Corporation.
# Licensed under the MIT license.
GOCMD=go
GOBUILD=$(GOCMD) build -v 
GOHOSTOS=$(strip $(shell $(GOCMD) env get GOHOSTOS))
GOTEST=GOOS=$(GOHOSTOS) $(GOCMD) test -v -coverprofile=coverage.out -covermode count -timeout 60m0s
PKGS=./...


TAG ?= $(shell git describe --tags)
COMMIT ?= $(shell git describe --always)
BUILD_DATE ?= $(shell date -u +%m/%d/%Y)

PKG := 

all: format library
 
.PHONY: vendor
vendor:
	GO111MODULE=on go mod tidy

library:
	GO111MODULE=on GOARCH=amd64 GOOS=windows $(GOBUILD) ./...

format:
	gofmt -s -w pkg

test:
	GOOS=windows GO111MODULE=on  GOARCH=amd64 go test -v ./go/...

prereq:
	git config --global url.ssh://git@github.com/.insteadOf https://github.com/
# if you see port already bound issues for etcd tests, consider limiting
# parallel execution using '-p 1'
unittest: prereq
	$(GOTEST) $(PKGS)