package prom

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Prom struct {
	registry *prometheus.Registry
}

func (p *Prom) Handler() http.Handler {
	return promhttp.InstrumentMetricHandler(
		p.registry, promhttp.HandlerFor(p.registry, promhttp.HandlerOpts{}),
	)
}

type Option func(*Prom)

func NewProm(options ...Option) *Prom {
	p := &Prom{
		registry: prometheus.NewRegistry(),
	}
	for _, option := range options {
		option(p)
	}
	return p
}

func WithSystem() Option {
	return func(prom *Prom) {
		prom.registry.MustRegister(
			Routine,
			Core,
			CPU,
			Mem,
			Disk,
			Load,
		)
	}
}

func WithDB(db *sql.DB) Option {
	Db = db
	return func(prom *Prom) {
		prom.registry.MustRegister(
			OpenConnections,
			Idle,
			InUse,
			WaitCount,
			WaitDuration,
		)
	}
}

type gaugeVecFuncCollector struct {
	desc *prometheus.Desc
	//TODO
	//gaugeVecFuncWithLabelValues []gaugeVecFuncWithLabelValues
	labelsDeduplicatedMap map[string]bool
}

// NewGaugeVecFunc
func NewGaugeVecFunc(opts prometheus.GaugeOpts, labelNames []string) *gaugeVecFuncCollector {
	return &gaugeVecFuncCollector{
		desc: prometheus.NewDesc(
			prometheus.BuildFQName(opts.Namespace, opts.Subsystem, opts.Name),
			opts.Help,
			labelNames,
			opts.ConstLabels,
		),
		labelsDeduplicatedMap: make(map[string]bool),
	}
}

// Describe
func (dc *gaugeVecFuncCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- dc.desc
}

// Collect
func (dc *gaugeVecFuncCollector) Collect(ch chan<- prometheus.Metric) {
	//TODO
	//for _, v := range dc.gaugeVecFuncWithLabelValues {
	//	ch <- prometheus.MustNewConstMetric(dc.desc, prometheus.GaugeValue, v.gaugeVecFunc(), v.labelValues...)
	//}
}

// RegisterGaugeVecFunc
// 同一组 labelValues 只能注册一次
func (dc *gaugeVecFuncCollector) RegisterGaugeVecFunc(labelValues []string, gaugeVecFunc func() float64) (err error) {
	// prometheus 每次允许收集一次 labelValues 相同的 metric
	deduplicateKey := strings.Join(labelValues, "")
	if dc.labelsDeduplicatedMap[deduplicateKey] {
		return fmt.Errorf("labelValues func already registered, labelValues:%v", labelValues)
	}
	dc.labelsDeduplicatedMap[deduplicateKey] = true
	//TODO
	//handlePanicGaugeVecFunc := func() float64 {
	//	if rec := recover(); rec != nil {
	//		const size = 10 * 1024
	//		buf := make([]byte, size)
	//		buf = buf[:runtime.Stack(buf, false)]
	//	}
	//	return gaugeVecFunc()
	//}
	//dc.gaugeVecFuncWithLabelValues = append(dc.gaugeVecFuncWithLabelValues, gaugeVecFuncWithLabelValues{
	//	gaugeVecFunc: handlePanicGaugeVecFunc,
	//	labelValues:  labelValues,
	//})
	return nil
}
