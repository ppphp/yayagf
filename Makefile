all: build test

export GOPRIVATE:=gitlab.papegames.com/*
export GOPROXY:=https://goproxy.io

build:
	go build

test:
	go test ...

update:
	go get -u

install:
	go install .
