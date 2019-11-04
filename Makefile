SHELL:=/bin/bash

all: build

test:
	go test ./... -v

build:
	go build -o ghp

clean:
	rm -rf ghp

.PHONY: all test build clean
