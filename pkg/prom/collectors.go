package prom

import (
	"database/sql"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"runtime"
)

func CPU() *gaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace:   "system",
		Subsystem:   "cpu",
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
}

func Mem() *gaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace:   "system",
		Subsystem:   "mem",
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
		lvs = append(lvs, LV{Lbs: []string{"swap_total"}, V: float64(m2.Total)})
		return
	})
}

func Disk() *gaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace:   "system",
		Subsystem:   "disk",
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
}

func Load() *gaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace:   "system",
		Subsystem:   "load",
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
}

func GoRoutine() prometheus.GaugeFunc {
	return prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace: "runtime",
		Subsystem: "goroutine",
		Name:      "goroutine",
	}, func() float64 {
		return float64(runtime.NumGoroutine())
	})
}

func UrlTTL() *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   "outer",
		Subsystem:   "http",
		Name:        "http_handler_ttl",
		ConstLabels: map[string]string{},
	}, []string{"url", "method", "ret"})
}

func UrlConnection() *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace:   "outer",
		Subsystem:   "http",
		Name:        "http_connections",
		ConstLabels: map[string]string{},
	}, []string{"url", "method"})
}

// https://orangematter.solarwinds.com/2018/05/22/new-stats-exposed-in-gos-database-sql-package/

// The number of connections.
func DbConnection(Db *sql.DB) *gaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace:   "db",
		Subsystem:   "connection",
		Name:        "connection",
		ConstLabels: map[string]string{},
	}, []string{"type"}, func() []LV {
		lvs := []LV{
			// The number of connections.
			{[]string{"open"}, float64(Db.Stats().OpenConnections)},
			// The number of open connections that are currently idle.
			{[]string{"idle"}, float64(Db.Stats().Idle)},
			// The number of connections actively in-use.
			{[]string{"inuse"}, float64(Db.Stats().InUse)},
		}
		return lvs
	})
}

// The total number of times a goroutine has had to wait for a connection.
func DBWaitCount(Db *sql.DB) prometheus.GaugeFunc {
	return prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace:   "db",
		Subsystem:   "wait",
		Name:        "db_wait_count",
		ConstLabels: map[string]string{},
	}, func() float64 {
		if Db == nil {
			return 0
		}
		return float64(Db.Stats().WaitCount)
	})
}

// The cumulative amount of time goroutines have spent waiting for a connection.
func DBWaitDuration(Db *sql.DB) prometheus.GaugeFunc {
	return prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace:   "db",
		Subsystem:   "wait",
		Name:        "db_wait_duration",
		ConstLabels: map[string]string{},
	}, func() float64 {
		if Db == nil {
			return 0
		}
		return Db.Stats().WaitDuration.Seconds()
	})
}
