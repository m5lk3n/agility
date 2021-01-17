#!/bin/bash

IMAGE="lttl.dev/k8s-df:latest"
CONTAINER=$(docker ps -aqf "ancestor=${IMAGE}")

if [ "${CONTAINER}" != "" ]; then
  docker rm -f ${CONTAINER}
fi