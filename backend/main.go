package main

import (
	nodeexporter "lttl.dev/k8s-df/backend/nodeexporter"
)

func main() {
	nodeexporter.Start()
}
