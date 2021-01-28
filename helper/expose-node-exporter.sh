#!/bin/bash

export POD_NAME=$(kubectl get pods --namespace k8s-df -l "app=k8s-df" -o jsonpath="{.items[0].metadata.name}")
export WEB_APP_PORT=8089
echo "Exposing node-exporter on port ${WEB_APP_PORT} ..."
kubectl --namespace k8s-df port-forward ${POD_NAME} ${WEB_APP_PORT}:8080