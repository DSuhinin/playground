package helpers

import (
	"github.com/prometheus/client_golang/prometheus"
)

//
// ResetPrometheusMetrics makes reset of Prometheus Metrics.
//
func ResetPrometheusMetrics() {

	prometheus.DefaultRegisterer = prometheus.NewRegistry()
}
