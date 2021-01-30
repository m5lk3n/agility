package nodeexporter

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

// Start a node exporter
func Start() {
	const port = ":8080"
	const endpoint = "/metrics"

	http.Handle(endpoint, promhttp.Handler())
	log.Infof("Node exporter serving from http://localhost%s%s", port, endpoint)
	log.Fatal(http.ListenAndServe(port, nil))
}
