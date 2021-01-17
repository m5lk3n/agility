#!/bin/bash

NAMESPACE=k8s-df
NAME=k8s-df
IMAGE="lttl.dev/k8s-df:${IMAGE_VER}"

kubectl delete namespace ${NAMESPACE}

kubectl create namespace ${NAMESPACE}
kubectl apply -f k8s-df-cm.yaml --namespace=${NAMESPACE}
# add CPU and memory limits and requests:
kubectl create deployment --image=${IMAGE} ${NAME} --namespace=${NAMESPACE}
kubectl expose deployment ${NAME} --port=80 --name=${NAME} --namespace=${NAMESPACE}