#!/bin/bash

if [ "${IMAGE_VER}" == "" ]; then
  echo "export IMAGE_VER"
  exit 1
fi

IMAGE="lttl.dev/k8s-df:${IMAGE_VER}"
CONTAINER=$(docker ps -aqf "ancestor=${IMAGE}")

if [ "${CONTAINER}" != "" ]; then
  docker rm -f ${CONTAINER}
fi