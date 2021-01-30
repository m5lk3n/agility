package nodeexporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

// DeployedTotalMetric is a KPI which is set by the deployment-watcher upon added deployment
var DeployedTotalMetric = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "deployed_total",
		Help: "Amount of app deployments into the cluster namespace"},
	[]string{"namespace", "app"},
)

func init() {
	prometheus.MustRegister(DeployedTotalMetric)
}
