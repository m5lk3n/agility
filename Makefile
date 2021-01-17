
IMAGE = lttl.dev/k8s-df:${IMAGE_VER}

# default target
.PHONY: help
help:
	@echo "usage: make <target>"
	@echo
	@echo "  where <target> is one of the following"
	@echo
	@echo "    clean       to remove the previously built container image with version 'IMAGE_VER'"
	@echo "    bake        to bake a new container image as version 'IMAGE_VER'"
	@echo "    load        to load the newly built image into kind"
	@echo "    redeploy    to uninstall and to install the application with the loaded image"
	@echo "    all         to run all targets"
	@echo
	@echo "    help        to show this text"

.PHONY: clean
clean:
	docker rmi ${IMAGE}

.PHONY: bake
bake:
	$(MAKE) -C node-exporter build
	$(MAKE) -C deployments-watcher build
	docker build -t ${IMAGE} .

.PHONY: load
load:
	kind load docker-image ${IMAGE}
	docker exec -it kind-control-plane crictl images

.PHONY: redeploy
redeploy:
	$(MAKE) -C k8s all

.PHONY: all
all: clean bake load redeploy