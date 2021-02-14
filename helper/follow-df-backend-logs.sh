#!/bin/bash

export POD_NAME=$(kubectl get pods --namespace k8s-df -o jsonpath="{.items[0].metadata.name}")
kubectl --namespace k8s-df logs $POD_NAME --follow
