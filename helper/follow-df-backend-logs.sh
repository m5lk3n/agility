#!/bin/bash

export NAMESPACE=agility
export POD_NAME=$(kubectl get pods --namespace ${NAMESPACE} -o jsonpath="{.items[0].metadata.name}")
kubectl --namespace ${NAMESPACE} logs $POD_NAME --follow
