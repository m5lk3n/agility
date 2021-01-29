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
	go mod init lttl.dev/k8s-df/backend

.PHONY: get
get:
	go get k8s.io/client-go@v0.19.1
	go get github.com/elankath/logflag
	go get github.com/sirupsen/logrus
	go get github.com/prometheus/client_golang/prometheus/promhttp
	go get github.com/prometheus/client_golang/prometheus

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o df-backend

.PHONY: all
all: init get build