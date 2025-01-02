SHELL := /bin/bash
GO111MODULE := on
GO := go

export GOPRIVATE := github.ibm.com

FIND_FLAGS := -type f -name '*.go' -not -name '*_test.go'
SRC := $(shell find ./internal $(FIND_FLAGS))


bin/%: ./cmd/main.go | bin
	$(GO) build -o $@ ./$(<D)

bin/FileTransfer: $(SRC) go.mod

bin:
	mkdir -p bin

.PHONY: clean
clean:
	$(RM) -r bin

.PHONY: build
build: clean bin/FileTransfer