package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var counterMetric = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "counter_metric",
		Help: "Counts occurred in our cluster"},
	[]string{"app", "namespace"},
)

func init() {

	prometheus.MustRegister(counterMetric)

	counterMetric.WithLabelValues("myApp", "myNamespace").Add(1)
}
