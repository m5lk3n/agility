
IMAGE = lttl.dev/k8s-df:${IMAGE_VER}
VERSION_TXT = version.txt
HELM_RELEASE = magility

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
	@echo "    new         to run all target but clean"
	@echo
	@echo "    help        to show this text"

.PHONY: check
check:
	[ "${IMAGE_VER}" != "" ] || echo "IMAGE_VER has to be set"

.PHONY: clean
clean: check
	docker rmi ${IMAGE}
	docker exec -it kind-control-plane crictl rmi ${IMAGE}
	helm uninstall ${HELM_RELEASE}

.PHONY: bake
bake: check
	$(MAKE) -C df-backend build
	$(MAKE) -C df-frontend build
	date > ${VERSION_TXT}
	echo "${IMAGE_VER}" >> ${VERSION_TXT}
	docker build -t ${IMAGE} .

.PHONY: load
load: check
	kind load docker-image ${IMAGE}
	docker exec -it kind-control-plane crictl images

.PHONY: redeploy
redeploy: check
	helm install ${HELM_RELEASE} --set image.version=${IMAGE_VER} ./chart

.PHONY: new
new: bake load redeploy