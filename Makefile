#! /usr/bin/make
#
# Makefile for Access Request microservice
#
# Top-level targets:
# init: install development tools and install dependencies
# update: Updates vendor dependencies
# tools: Install/upgrade development tools
# all: builds all microservices
# test: Runs go test on all found packages
# containers: Wraps builded microservices into Docker containers

# Dependencies that are used in this project by the build&test process
DEPEND=github.com/smartystreets/goconvey

# Microservice name
export MS_NAME=mustrum

# Internal variables
MKDIR_P = mkdir -p

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
export CURRENT_DIR := $(patsubst %/,%,$(dir $(mkfile_path)))
export OUTDIR := $(CURRENT_DIR)/build

dir_target = $(OUTDIR)-$(wildcard $(OUTDIR))
dir_present = $(OUTDIR)-$(OUTDIR)
dir_absent = $(OUTDIR)-
VERSION_FILE := $(OUTDIR)/current_version

# Check presence of dependencies
SWAGGER_VERSION_REQUIRED=v0.19.0
GLIDE := $(shell command -v glide 2> /dev/null)
SWAGGER := $(shell command -v swagger 2> /dev/null)
GOCONVEY := $(shell command -v goconvey 2> /dev/null)
SWAGGERVERSION := $(shell swagger version 2> /dev/null | head -n 1 - | cut -d' ' -f2)
OS_FLAVOUR := $(shell uname | tr '[:upper:]' '[:lower:]')
OS_MARCH := $(shell uname -m)
ifeq (${OS_MARCH}, x86_64)
	OS_MARCH = amd64
else
	OS_MARCH = 386
endif

ifeq ($(origin GOBIN), undefined)
gobin=${HOME}/go/bin
else
gobin=${GOBIN}
endif

.PHONY: all init update containers

all: build
init: tools update

.PHONY: tools
tools:
ifndef GLIDE
	wget -q -O - https://glide.sh/get | sh 2>/dev/null
endif
ifndef SWAGGER
	wget -q -O "${gobin}/swagger" "https://github.com/go-swagger/go-swagger/releases/download/${SWAGGER_VERSION_REQUIRED}/swagger_${OS_FLAVOUR}_${OS_MARCH}"
	chmod a+x "${gobin}/swagger"
else
ifneq "${SWAGGERVERSION}" "${SWAGGER_VERSION_REQUIRED}"
	mv "${gobin}/swagger" "${gobin}/swagger-${SWAGGERVERSION}"
	wget -q -O "${gobin}/swagger" "https://github.com/go-swagger/go-swagger/releases/download/${SWAGGER_VERSION_REQUIRED}/swagger_${OS_FLAVOUR}_${OS_MARCH}"
	chmod a+x "${gobin}/swagger"
endif
endif
ifndef GOCONVEY
	go get -v $(DEPEND)
endif

update:
	@echo "Fetch project dependencies..."
	glide install

.PHONY: generate
generate:
	@if [ -d server/restapi/operations ]; then find server/restapi/operations -name '*.go' -delete; fi
	@if [ -d server/models ]; then find server/models -name '*.go' -delete; fi
	swagger generate server --target server --name $(MS_NAME) \
		--spec server/swagger.yaml --flag-strategy pflag

.PHONY: build
build:
	@echo "Build project binary..."
	$(eval BUILDER = $(shell git config user.name | tr ' ' '.'))
	$(eval GPATH = $(shell go env | grep GOPATH | cut -d'=' -f2 | tr -d \"))
	CGO_ENABLED=0 go build -o build/$(MS_NAME) cmd/$(MS_NAME)-server/main.go

.PHONY: test
test:
	go test $(shell go list ./... | egrep -v "\/cmd\/|\/models|\/restapi?") -count=1

.PHONY: go-test
go-test: test

.PHONY: go-build
go-build: build

.PHONY: go-get
go-get: update

.PHONY: clean
clean:
	rm -rf $(OUTDIR)/$(MS_NAME)*

