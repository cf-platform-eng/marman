SHELL = /bin/bash
GO-VER = go1.25

default: build

# #### GO Binary Management ####
deps-go-binary:
	echo "Expect: $(GO-VER)" && \
		echo "Actual: $$(go version)" && \
	 	go version | grep $(GO-VER) > /dev/null

# #### CLEAN ####
clean:
	rm -rf build/*
	go clean --modcache

# #### DEPS ####
deps-modules:
	go mod download

deps: deps-modules

# #### BUILD ####
SRC = $(shell find . -name "*.go" | grep -v "_test\." )
VERSION := $(or $(VERSION), dev)
LDFLAGS="-X github.com/cf-platform-eng/marman/version.Version=$(VERSION)"

build/marman: $(SRC) deps
	go build -o build/marman -ldflags ${LDFLAGS} ./cmd/marman/main.go

build: build/marman

build-all: build-linux build-darwin

build-linux: build/marman-linux

build/marman-linux: $(SRC) deps
	GOARCH=amd64 GOOS=linux go build -o build/marman-linux -ldflags ${LDFLAGS} ./cmd/marman/main.go

build-darwin: build/marman-darwin

build/marman-darwin: $(SRC) deps
	GOARCH=amd64 GOOS=darwin go build -o build/marman-darwin -ldflags ${LDFLAGS} ./cmd/marman/main.go

build-image: build/marman-linux
	docker build --tag cfplatformeng/marman:${VERSION} --file Dockerfile .

# #### TESTS ####
test-units: deps lint
	go tool ginkgo -r --skipPackage features .

test-features: deps
	go tool ginkgo -r --tags=feature features

test: test-units test-features

lint: deps

.PHONY: set-pipeline
set-pipeline: ci/pipeline.yaml
	fly -t ppe-isv set-pipeline -p marman -c ci/pipeline.yaml
