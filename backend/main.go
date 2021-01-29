package main

import (
	deploymentswatcher "lttl.dev/k8s-df/backend/deploymentswatcher"
	nodeexporter "lttl.dev/k8s-df/backend/nodeexporter"
)

func main() {
	go deploymentswatcher.Start()
	nodeexporter.Start()
}
