package nodeexporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var deployedTotalMetric = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "deployed_total",
		Help: "Amount of app deployments into the cluster namespace"},
	[]string{"namespace", "app"},
)

func init() {
	prometheus.MustRegister(deployedTotalMetric)
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
