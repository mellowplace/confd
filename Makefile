.PHONY: test build
GOPATH=$(shell pwd)

deps:
	bin/godep restore
test:
	bin/godep go test
build:
	go build -o bin/confd .

build-docker:
	docker build -t confd_builder -f Dockerfile.build.alpine .
	docker run -ti --rm -v $(shell pwd):/app confd_builder ./build
test-docker:
	docker run -ti --rm -v $(shell pwd):/app confd_builder ./test
