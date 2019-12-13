all: build test

build:
	go build

test:
	go test ...

update:
	go get -u

install:
	go install .
