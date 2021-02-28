#!/bin/bash

for i in {1..50}
do
  deployment="hello-echo-$((i % 2))"
  date
  echo "deployment: $i $deployment"
  kubectl delete deployment $deployment
  sleep 1
  kubectl create deployment $deployment --image=k8s.gcr.io/echoserver:1.10
  sleep 1 
done
