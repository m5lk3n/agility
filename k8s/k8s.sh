#!/bin/bash

NAMESPACE=k8s-df
kubectl create namespace ${NAMESPACE}
kubectl apply -f k8s-df-cm.yaml --namespace=${NAMESPACE}