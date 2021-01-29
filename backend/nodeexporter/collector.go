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

	deployedTotalMetric.WithLabelValues("myNamespace", "myApp").Inc()
}
