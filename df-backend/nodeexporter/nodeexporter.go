package nodeexporter

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var deployedTotalMetric = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "deployed_total",
		Help: "Amount of app deployments into the cluster namespace"},
	[]string{"namespace", "app"},
)

var handler http.Handler

func init() {
	// https://stackoverflow.com/questions/35117993/how-to-disable-go-collector-metrics-in-prometheus-client-golang:
	// "go get rid of any additional metrics
	//  we have to expose our metrics with a custom registry"
	r := prometheus.NewRegistry()
	r.MustRegister(deployedTotalMetric)
	handler = promhttp.HandlerFor(r, promhttp.HandlerOpts{})

	// the above renders the following flag useless:
	// https://github.com/prometheus/node_exporter/pull/1148
	// --web.disable-exporter-metrics
}

// Deployment is the container for the metric's details
type Deployment struct {
	Name      string
	Namespace string
}

// Add a deployment incrementing corresponding counter
func Add(deployment Deployment) {
	deployedTotalMetric.WithLabelValues(deployment.Namespace, deployment.Name).Inc()
}

// Start a node exporter
func Start() {
	const port = ":8080"
	const endpoint = "/metrics"

	http.Handle(endpoint, handler)
	log.Infof("node exporter serving from http://localhost%s%s\n", port, endpoint)
	log.Fatal(http.ListenAndServe(port, nil))
}
