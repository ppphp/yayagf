package prom

import (
	"fmt"
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

var CPU = NewGaugeVecFunc(prometheus.GaugeOpts{
	Namespace:   "yayagf",
	Subsystem:   "system",
	Name:        "cpu",
	ConstLabels: map[string]string{},
}, []string{"num"}, func() []LV {
	fs, err := cpu.Percent(0, true)
	if err != nil {
		return nil
	}
	lvs := []LV{}
	for k, v := range fs {
		lvs = append(lvs, LV{Lbs: []string{fmt.Sprint(k)}, V: v})
	}
	return lvs
})

var Mem = NewGaugeVecFunc(prometheus.GaugeOpts{
	Namespace:   "yayagf",
	Subsystem:   "system",
	Name:        "mem",
	ConstLabels: map[string]string{},
}, []string{"cat"}, func() (lvs []LV) {
	lvs = []LV{}
	m, err := mem.VirtualMemory()
	if err != nil {
		return
	}
	lvs = append(lvs, LV{Lbs: []string{"virtual_free"}, V: float64(m.Free)})
	lvs = append(lvs, LV{Lbs: []string{"virtual_used"}, V: float64(m.Used)})
	lvs = append(lvs, LV{Lbs: []string{"virtual_total"}, V: float64(m.Total)})
	lvs = append(lvs, LV{Lbs: []string{"virtual_active"}, V: float64(m.Active)})
	m2, err := mem.SwapMemory()
	if err != nil {
		return
	}
	lvs = append(lvs, LV{Lbs: []string{"swap_free"}, V: float64(m2.Free)})
	lvs = append(lvs, LV{Lbs: []string{"swap_used"}, V: float64(m2.Used)})
	lvs = append(lvs, LV{Lbs: []string{"swap_used"}, V: float64(m2.Total)})
	return
})

var Disk = NewGaugeVecFunc(prometheus.GaugeOpts{
	Namespace:   "yayagf",
	Subsystem:   "system",
	Name:        "disk",
	ConstLabels: map[string]string{},
}, []string{"path"}, func() []LV {
	ps, err := disk.Partitions(false)
	if err != nil {
		return nil
	}
	lvs := []LV{}
	for _, v := range ps {
		d, err := disk.Usage(v.Mountpoint)
		if err == nil {
			lvs = append(lvs, LV{Lbs: []string{v.Mountpoint}, V: float64(d.Free)})
		}
	}
	return lvs
})

var Load = NewGaugeVecFunc(prometheus.GaugeOpts{
	Namespace:   "yayagf",
	Subsystem:   "system",
	Name:        "load",
	ConstLabels: map[string]string{},
}, []string{"cat"}, func() (lvs []LV) {
	lvs = []LV{}
	a, err := load.Avg()
	if err != nil {
		return
	}
	lvs = append(lvs, LV{Lbs: []string{"avg_1"}, V: float64(a.Load1)})
	lvs = append(lvs, LV{Lbs: []string{"avg_5"}, V: float64(a.Load5)})
	lvs = append(lvs, LV{Lbs: []string{"avg_15"}, V: float64(a.Load15)})
	m, err := load.Misc()
	if err != nil {
		return
	}
	lvs = append(lvs, LV{Lbs: []string{"misc_total"}, V: float64(m.ProcsTotal)})
	lvs = append(lvs, LV{Lbs: []string{"misc_running"}, V: float64(m.ProcsRunning)})
	lvs = append(lvs, LV{Lbs: []string{"misc_blocked"}, V: float64(m.ProcsBlocked)})
	return lvs
})
