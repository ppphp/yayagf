package prom

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func PromHandler() http.Handler {
	reg := prometheus.NewRegistry()

	return promhttp.InstrumentMetricHandler(
		reg, promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}),
	)
}
