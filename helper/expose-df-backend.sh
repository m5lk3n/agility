#!/bin/bash

export NAMESPACE=agility
export POD_NAME=$(kubectl get pods --namespace ${NAMESPACE} -l "app=k8s-df" -o jsonpath="{.items[0].metadata.name}")
export WEB_APP_PORT=8089
echo "Exposing df-backend (node-exporter) on port ${WEB_APP_PORT} ..."
kubectl --namespace ${NAMESPACE} port-forward ${POD_NAME} ${WEB_APP_PORT}:8080