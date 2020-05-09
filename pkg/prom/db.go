package prom

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
)

var Db *sql.DB

var OpenConnections = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
	Namespace:   "yayagf",
	Subsystem:   "db",
	Name:        "open_connections",
	ConstLabels: map[string]string{},
}, func() float64 {
	if Db == nil {
		return 0
	}
	return float64(Db.Stats().OpenConnections)
})

var Idle = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
	Namespace:   "yayagf",
	Subsystem:   "db",
	Name:        "idle",
	ConstLabels: map[string]string{},
}, func() float64 {
	if Db == nil {
		return 0
	}
	return float64(Db.Stats().Idle)
})

var InUse = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
	Namespace:   "yayagf",
	Subsystem:   "db",
	Name:        "inuse",
	ConstLabels: map[string]string{},
}, func() float64 {
	if Db == nil {
		return 0
	}
	return float64(Db.Stats().InUse)
})

var WaitCount = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
	Namespace:   "yayagf",
	Subsystem:   "db",
	Name:        "waitcount",
	ConstLabels: map[string]string{},
}, func() float64 {
	if Db == nil {
		return 0
	}
	return float64(Db.Stats().WaitCount)
})

var WaitDuration = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
	Namespace:   "yayagf",
	Subsystem:   "db",
	Name:        "waitduration",
	ConstLabels: map[string]string{},
}, func() float64 {
	if Db == nil {
		return 0
	}
	return float64(Db.Stats().WaitDuration.Seconds())
})


