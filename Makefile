SHELL = /bin/bash
GO-VER = go1.12

default: build

# #### GO Binary Management ####
deps-go-binary:
	echo "Expect: $(GO-VER)" && \
		echo "Actual: $$(go version)" && \
	 	go version | grep $(GO-VER) > /dev/null


HAS_GO_IMPORTS := $(shell command -v goimports;)

deps-goimports: deps-go-binary
ifndef HAS_GO_IMPORTS
	go get -u golang.org/x/tools/cmd/goimports
endif

# #### CLEAN ####
clean: deps-go-binary
	rm -rf build/*
	go clean --modcache


# #### DEPS ####

deps: deps-goimports deps-go-binary
	go mod download

# #### BUILD ####
SRC = $(shell find . -name "*.go" | grep -v "_test\." )

VERSION := $(or $(VERSION), "dev")

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

test: deps lint
	ginkgo -r .

lint: deps-goimports
	git ls-files | grep '.go$$' | xargs goimports -l -w
