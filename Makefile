.PHONY: all clean deps lint test build ui logmgr

COMMIT ?= $(shell git rev-parse --short HEAD)
LDFLAGS = -X "main.buildCommit=$(COMMIT)"
PACKAGES = $(shell go list ./... | grep -v /vendor/)

all: deps build

clean:
	if [ -f logmgr ] ; then rm -f logmgr ; fi

deps:
	go get ./...
	npm i

lint:
	go fmt $(PACKAGES)
	go vet $(PACKAGES)

test:
	@for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

build: ui logmgr

ui:
	node_modules/.bin/webpack

logmgr: $(wildcard *.go)
	CGO_ENABLED=0 go build -ldflags '-s -w $(LDFLAGS)' -o logmgr
