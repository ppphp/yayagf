package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

type LV struct {
	Lbs []string
	V   float64
}

type gaugeVecFuncCollector struct {
	desc                  *prometheus.Desc
	labelsDeduplicatedMap map[string]bool
	funcs                 func() []LV
}

// NewGaugeVecFunc
func NewGaugeVecFunc(opts prometheus.GaugeOpts, labelNames []string, funcs func() []LV) *gaugeVecFuncCollector {
	return &gaugeVecFuncCollector{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(opts.Namespace, opts.Subsystem, opts.Name),
			opts.Help,
			labelNames,
			opts.ConstLabels,
		),
		labelsDeduplicatedMap: make(map[string]bool),
		funcs:                 funcs,
	}
}

// Describe
func (dc *gaugeVecFuncCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- dc.desc
}

// Collect
func (dc *gaugeVecFuncCollector) Collect(ch chan<- prometheus.Metric) {
	for _, v := range dc.funcs() {
		ch <- prometheus.MustNewConstMetric(dc.desc, prometheus.GaugeValue, v.V, v.Lbs...)
	}
}
