#!/bin/bash

# deploys a dummy app; alternating app-1, app-2

for i in {1..50}
do
  deployment="app-$((i % 2))"
  date
  echo "app deployment: $i $deployment"
  kubectl delete deployment $deployment
  sleep 1
  kubectl create deployment $deployment --image=k8s.gcr.io/echoserver:1.10
  sleep 1 
done
