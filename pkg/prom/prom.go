package prom

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
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
