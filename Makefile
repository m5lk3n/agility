DEFAULT_IMAGE_VER = "0.1.0"
IMAGE_VER ?= ${DEFAULT_IMAGE_VER} # can be overridden by ENV param 
IMAGE = lttl.dev/agility-df:${IMAGE_VER}
VERSION_TXT = version.txt
HELM_RELEASE = agility
NAMESPACE = agility

# default target
.PHONY: help
help:
	@echo "builds, (un)deploys, ... '${HELM_RELEASE}' using docker, kind, and helm"
	@echo
	@echo "hint: use 'export IMAGE_VER=major.minor.patch' to overwrite the default '${DEFAULT_IMAGE_VER}'"
	@echo
	@echo "usage: make <target>"
	@echo
	@echo "  where <target> is one of the following"
	@echo
	@echo "    clean       to remove the container image '${IMAGE}' (incl. the version loaded to kind)"
	@echo "    uninstall   to undeploy the application '${HELM_RELEASE}' currently installed on the kind K8s cluster"
	@echo "    bake        to bake a new container image as '${IMAGE}'"
	@echo "    load        to load the container image ('${IMAGE}') into kind"
	@echo "    install     to deploy the application with the loaded image ('${IMAGE}')"
	@echo "    new         to run all targets but clean and uninstall"
	@echo
	@echo "    help        to show this text"

# checks existence of required tool stack
.PHONY: check
check:
	kind --version > /dev/null
	helm version > /dev/null
	docker --version > /dev/null

.PHONY: clean
clean: check
	docker rmi ${IMAGE}
	docker exec -it kind-control-plane crictl rmi ${IMAGE}

.PHONY: uninstall
uninstall:
	helm uninstall --namespace ${NAMESPACE} ${HELM_RELEASE}

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

.PHONY: install
install: check
	helm install --namespace ${NAMESPACE} --create-namespace --set image.version=${IMAGE_VER} ${HELM_RELEASE} ./chart

.PHONY: new
new: bake load install