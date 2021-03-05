#!/bin/bash

# deploys the dummy app-1

for i in {1..6}
do
  deployment="app-1"
  date
  echo "app deployment: $i $deployment"
  kubectl delete deployment $deployment
  sleep 1
  kubectl create deployment $deployment --image=k8s.gcr.io/echoserver:1.10
  sleep 1 
done
