#!/bin/bash

IMAGE="lttl.dev/k8s-df:${IMAGE_VER}"

docker run --name k8s-df --rm -i -t ${IMAGE} sh