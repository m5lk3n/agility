#!/bin/bash

IMAGE="lttl.dev/k8s-df:latest"

docker run --name k8s-df --rm -i -t ${IMAGE} sh