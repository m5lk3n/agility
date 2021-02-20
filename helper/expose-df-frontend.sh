#!/bin/bash

export APP_NAME=agility
export NAMESPACE=agility
export POD_NAME=$(kubectl get pods --namespace ${NAMESPACE} -l "app=${APP_NAME}-df" -o jsonpath="{.items[0].metadata.name}")
export WEB_APP_PORT=8088
echo "Exposing df-frontend (web-app) on port ${WEB_APP_PORT} ..."
kubectl --namespace ${NAMESPACE} port-forward ${POD_NAME} ${WEB_APP_PORT}:80