package prom

import "github.com/prometheus/client_golang/prometheus"

func NewUrlCounter() *prometheus.CounterVec {
	return prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace:   "yayagf",
		Subsystem:   "url",
		Name:        "counter",
		ConstLabels: map[string]string{},
	}, []string{"url", "method", "ret"})
}

func NewUrlTTLHist() *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   "yayagf",
		Subsystem:   "ttl",
		Name:        "hist",
		ConstLabels: map[string]string{},
	}, []string{"url", "method", "ret"})
}
