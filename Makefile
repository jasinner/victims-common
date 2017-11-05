CODE_DIRS := db/ log/ types/

.PHONY: help deps clean gofmt golint check

default: help

help:
	@echo "Targets:"
	@echo "	deps: Install dependencies with govendor"
	@echo "	build: Build the library"
	@echo "	clean: cleans up and removes built files"
	@echo "	gofmt: Run gofmt against the code"
	@echo "	golint: Run golint against the code"
	@echo "	check: Run all checks against the code"

deps:
	go get github.com/kardianos/govendor
	govendor sync

build:
	govendor build ./...

clean:
	go clean
	rm -f victims-api

gofmt:
	gofmt -l ${CODE_DIRS}

golint:
	go get github.com/golang/lint/golint
	golint ${CODE_DIRS}

check: gofmt golint
