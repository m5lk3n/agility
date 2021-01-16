# default target
.PHONY: help
help:
	@echo "usage: make <target>"
	@echo
	@echo "  where <target> is one of the following"
	@echo
	@echo "    clean       to remove the previously built latest container image"
	@echo "    bake        to bake a new container image as latest"
	@echo "    all         to run all targets"
	@echo
	@echo "    help        to show this text"

.PHONY: clean
clean:
	docker rmi lttl.dev/k8s-df:latest

.PHONY: bake
bake:
	docker build -t lttl.dev/k8s-df .

.PHONY: all
all: clean bake
