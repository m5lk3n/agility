package main

import (
	deploymentswatcher "lttl.dev/k8s-df/df-backend/deploymentswatcher"
	nodeexporter "lttl.dev/k8s-df/df-backend/nodeexporter"
)

func main() {
	go deploymentswatcher.Start() // starts listener in background, informs node-exporter
	nodeexporter.Start()          // blocking call, server start in foreground
}
