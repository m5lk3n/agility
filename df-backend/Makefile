# default target
.PHONY: help
help:
	@echo "usage: make <target>"
	@echo
	@echo "  where <target> is one of the following"
	@echo
	@echo "    clean       to delete the go module"
	@echo "    init        to initialize the module"
	@echo "    get         to fetch all package dependencies"
	@echo "    build       to compile binary for linux amd64 architecture"
	@echo "    all         to run all targets"
	@echo
	@echo "    help        to show this text"

.PHONY: clean
clean:
	rm -f go.mod

.PHONY: init
init:
	go mod init lttl.dev/agility/df-backend

.PHONY: get
get:
	go get github.com/gogo/protobuf@v1.3.2
	go get k8s.io/client-go/plugin/pkg/client/auth/exec@v0.19.1
	go get k8s.io/client-go/discovery@v0.19.1
	go get k8s.io/client-go@v0.19.1
	go get k8s.io/client-go/rest@v0.19.1
	go get k8s.io/client-go/transport@v0.19.1
	go get github.com/elankath/logflag
	go get github.com/sirupsen/logrus
	go get github.com/prometheus/client_golang
	go get github.com/prometheus/client_golang/prometheus/promhttp
	go get github.com/prometheus/client_golang/prometheus

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o df-backend

.PHONY: all
all: clean init get build