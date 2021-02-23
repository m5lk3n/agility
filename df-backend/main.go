package main

import (
	deploymentswatcher "lttl.dev/agility/df-backend/deploymentswatcher"
	nodeexporter "lttl.dev/agility/df-backend/nodeexporter"
)

func main() {
	go deploymentswatcher.Start() // starts listener in background, informs node-exporter
	nodeexporter.Start()          // blocking call, server start in foreground
}
