all: build test fmt

.PHONY: fmt build test update install

export GOPRIVATE:=gitlab.papegames.com/*
export GOPROXY:=https://goproxy.io

Files=$(shell fd -tf -E testdata -e go)

fmt:
	goimports -w ${Files}

build:
	go build

test:
	go test ${Files}

update:
	go get -u

install:
	go install .
