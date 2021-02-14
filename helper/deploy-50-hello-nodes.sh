#!/bin/bash

for i in {1..50}
do
  echo "deployment: $i"
  kubectl delete deployment hello-node
  sleep 1
  kubectl create deployment hello-node --image=k8s.gcr.io/echoserver:1.4
  sleep 2 
done
