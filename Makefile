# there is dev, test and prod mode of the code
# when in dev mode, deps are local
# when in test mode, deps are vendor, and update deps by a certain command by hand every month.
# when in prod mode, deps are all vendor

all: build test

build:
	go build ./cmd/yayagf

test:
	go test ...

update:
	go get -u
