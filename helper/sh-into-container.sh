#!/bin/bash

if [ "${IMAGE_VER}" == "" ]; then
  echo "export IMAGE_VER"
  exit 1
fi

IMAGE="lttl.dev/k8s-df:${IMAGE_VER}"

docker run --name k8s-df --rm -i -t ${IMAGE} sh