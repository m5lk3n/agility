NAMESPACE = k8s-df
NAME = k8s-df
IMAGE = "lttl.dev/k8s-df:${IMAGE_VER}"

# default target
.PHONY: help
help:
	@echo "usage: make <target>"
	@echo
	@echo "  where <target> is one of the following"
	@echo
	@echo "    uninstall   to delete the ${NAMESPACE} namespace"
	@echo "    install     to create the ${NAMESPACE} namespace, and install the application in it"
	@echo "    all         to run all targets"
	@echo
	@echo "    help        to show this text"

.PHONY: uninstall
uninstall:
	kubectl delete namespace ${NAMESPACE}

.PHONY: install
install:
	kubectl create namespace ${NAMESPACE}
	kubectl apply -f . --namespace=${NAMESPACE}
	# add CPU and memory limits and requests to the creation:
	kubectl create deployment --image=${IMAGE} ${NAME} --namespace=${NAMESPACE}
	kubectl set serviceaccount deployment ${NAME} k8s-df-serviceaccount --namespace=${NAMESPACE}
	kubectl expose deployment ${NAME} --port=80 --name=${NAME} --namespace=${NAMESPACE}

.PHONY: all
all: uninstall install