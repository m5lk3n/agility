# default target
.PHONY: help
help:
	@echo "usage: make <target>"
	@echo
	@echo "  where <target> is one of the following"
	@echo
	@echo "    init        to initialize the module"
	@echo "    get         to fetch all package dependencies"
	@echo "    build       to compile binary for linux amd64 architecture"
	@echo "    all         to run all targets"
	@echo
	@echo "    help        to show this text"

.PHONY: init
init:
	go mod init

.PHONY: get
get:
	go get github.com/gin-gonic/contrib/static
	go get github.com/gin-gonic/gin

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o web-app

.PHONY: all
all: init get build