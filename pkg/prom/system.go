package prom

import (
	"runtime"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

var Routine = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
	Namespace: "yayagf",
	Subsystem: "system",
	Name:      "runtime",
}, func() float64 {
	return float64(runtime.NumGoroutine())
})

var Core = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
	Namespace: "yayagf",
	Subsystem: "system",
	Name:      "core",
}, func() float64 {
	return float64(runtime.NumCPU())
})

var CPU = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
	Namespace:   "yayagf",
	Subsystem:   "system",
	Name:        "cpu",
	ConstLabels: map[string]string{},
}, func() float64 {
	fs, _ := cpu.Percent(0, false)
	if len(fs) > 0 {
		return fs[0]
	} else {
		return 0
	}
})

var Mem = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
	Namespace:   "yayagf",
	Subsystem:   "system",
	Name:        "mem",
	ConstLabels: map[string]string{},
}, func() float64 {
	m, _ := mem.VirtualMemory()
	if m != nil {
		return m.UsedPercent
	} else {
		return 0
	}
})

var Disk = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
	Namespace:   "yayagf",
	Subsystem:   "system",
	Name:        "disk",
	ConstLabels: map[string]string{},
}, func() float64 {
	d, _ := disk.Usage("/")
	if d != nil {
		return d.UsedPercent
	} else {
		return 0
	}
})

var Load = prometheus.NewGaugeFunc(prometheus.GaugeOpts{
	Namespace:   "yayagf",
	Subsystem:   "system",
	Name:        "load",
	ConstLabels: map[string]string{},
}, func() float64 {
	a, _ := load.Avg()
	if a != nil {
		return a.Load15
	} else {
		return 0
	}
})
