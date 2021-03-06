.PHONY: test deps docs
.EXPORT_ALL_VARIABLES:

GO111MODULE ?= on
LOCALS      := $(shell find . -type f -name '*.go')

all: deps test build docs

deps:
	go get ./...
	-go mod tidy
	@go list github.com/mjibson/esc || go get github.com/mjibson/esc/...
	go generate -x ./...

fmt:
	gofmt -w $(LOCALS)
	go vet ./...

test:
	go test ./...

build: fmt
	test -d diecast && go build -i -o bin/diecast cmd/diecast/main.go
	which diecast && cp -v bin/diecast `which diecast` || true

docs:
	cd docs && make

package:
	-rm -rf pkg
	mkdir -p pkg/usr/bin
	cp bin/diecast pkg/usr/bin/diecast
	fpm \
		--input-type  dir \
		--output-type deb \
		--deb-user    root \
		--deb-group   root \
		--name        diecast \
		--version     `./pkg/usr/bin/diecast -v | cut -d' ' -f3` \
		-C            pkg
