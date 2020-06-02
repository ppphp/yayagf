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

func SysCPU(project string) *gaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace:   project,
		Name:        "system_cpu",
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

func SysMem(project string) *gaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace:   project,
		Name:        "system_mem",
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

func SysDisk(project string) *gaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace:   project,
		Name:        "system_disk",
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

func SysLoad(project string) *gaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace:   project,
		Name:        "system_load",
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

func GoRoutine(project string) prometheus.GaugeFunc {
	return prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace:   project,
		Name:      "goroutines",
	}, func() float64 {
		return float64(runtime.NumGoroutine())
	})
}

// 注释 https://blog.csdn.net/m0_38132420/article/details/71699815
func GoMem(project string) *gaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace:   project,
		Name:      "process_mem",
	}, []string{"cat"}, func() (lvs []LV) {
		lvs = []LV{}
		memstat := &runtime.MemStats{}
		runtime.ReadMemStats(memstat)
		lvs = append(lvs, LV{Lbs: []string{"sys"}, V: float64(memstat.Sys)})                    //服务现在系统使用的内存
		lvs = append(lvs, LV{Lbs: []string{"alloc"}, V: float64(memstat.Alloc)})                //golang语言框架堆空间分配的字节数
		lvs = append(lvs, LV{Lbs: []string{"total_alloc"}, V: float64(memstat.TotalAlloc)})     //从服务开始运行至今分配器为分配的堆空间总 和，只有增加，释放的时候不减少
		lvs = append(lvs, LV{Lbs: []string{"frees"}, V: float64(memstat.Frees)})                //服务回收的heap objects的字节数
		lvs = append(lvs, LV{Lbs: []string{"heap_alloc"}, V: float64(memstat.HeapAlloc)})       //服务分配的堆内存字节数
		lvs = append(lvs, LV{Lbs: []string{"heap_sys"}, V: float64(memstat.HeapSys)})           //系统分配的作为运行栈的内存
		lvs = append(lvs, LV{Lbs: []string{"heap_idle"}, V: float64(memstat.HeapIdle)})         //申请但是未分配的堆内存或者回收了的堆内存（空闲）字节数
		lvs = append(lvs, LV{Lbs: []string{"heap_inuse"}, V: float64(memstat.HeapInuse)})       //正在使用的堆内存字节数
		lvs = append(lvs, LV{Lbs: []string{"heap_released"}, V: float64(memstat.HeapReleased)}) //返回给OS的堆内存，类似C/C++中的free。
		lvs = append(lvs, LV{Lbs: []string{"stack_inuse"}, V: float64(memstat.StackInuse)})     //正在使用的栈字节数
		lvs = append(lvs, LV{Lbs: []string{"stack_sys"}, V: float64(memstat.StackSys)})         //系统分配的作为运行栈的内存
		lvs = append(lvs, LV{Lbs: []string{"mspan_inuse"}, V: float64(memstat.MSpanInuse)})     // mspan正在使用的
		lvs = append(lvs, LV{Lbs: []string{"mspan_sys"}, V: float64(memstat.MSpanSys)})         // mspan分配的的
		lvs = append(lvs, LV{Lbs: []string{"mcache_inuse"}, V: float64(memstat.MCacheInuse)})   // mcache结构体申请的字节数
		lvs = append(lvs, LV{Lbs: []string{"mcache_sys"}, V: float64(memstat.MCacheSys)})       // MCache分配的的
		lvs = append(lvs, LV{Lbs: []string{"buck_hash_sys"}, V: float64(memstat.BuckHashSys)})  // 这是啥？
		lvs = append(lvs, LV{Lbs: []string{"gc_sys"}, V: float64(memstat.GCSys)})               // GC分配的的
		lvs = append(lvs, LV{Lbs: []string{"other_sys"}, V: float64(memstat.OtherSys)})         //golang系统架构占用的额外空间
		return lvs
	})
}

func UrlTTL(project string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   project,
		Name:        "http_handler_ttl",
		ConstLabels: map[string]string{},
	}, []string{"url", "method", "ret"})
}

func UrlConnection(project string) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace:   project,
		Name:        "http_connections",
		ConstLabels: map[string]string{},
	}, []string{"url", "method"})
}

// https://orangematter.solarwinds.com/2018/05/22/new-stats-exposed-in-gos-database-sql-package/

// The number of connections.
func DbConnection(project string, dbname string, client *sql.DB) *gaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace:   project,
		Subsystem:   dbname,
		Name:        "db_connection",
		ConstLabels: map[string]string{},
	}, []string{"type"}, func() []LV {
		lvs := []LV{
			// The number of connections.
			{[]string{"open"}, float64(client.Stats().OpenConnections)},
			// The number of open connections that are currently idle.
			{[]string{"idle"}, float64(client.Stats().Idle)},
			// The number of connections actively in-use.
			{[]string{"inuse"}, float64(client.Stats().InUse)},
		}
		return lvs
	})
}

// The total number of times a goroutine has had to wait for a connection.
func DBWaitCount(project string, dbname string, client *sql.DB) prometheus.GaugeFunc {
	return prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace:   project,
		Subsystem:   dbname,
		Name:        "db_wait_count",
		ConstLabels: map[string]string{},
	}, func() float64 {
		if client == nil {
			return 0
		}
		return float64(client.Stats().WaitCount)
	})
}

// The cumulative amount of time goroutines have spent waiting for a connection.
func DBWaitDuration(project string, dbname string, client *sql.DB) prometheus.GaugeFunc {
	return prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace:   project,
		Subsystem:   dbname,
		Name:        "db_wait_duration",
		ConstLabels: map[string]string{},
	}, func() float64 {
		if client == nil {
			return 0
		}
		return client.Stats().WaitDuration.Seconds()
	})
}

func CallHTTPConnection(project string, name string) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace:   project,
		Subsystem:   name,
		Name:        "http_connections",
		ConstLabels: map[string]string{},
	}, []string{"url", "method"})
}

func CallHTTPTTL(project string, name string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   project,
		Subsystem:   name,
		Name:        "http_ttl",
		ConstLabels: map[string]string{},
	}, []string{"url", "method", "code", "ret"})
}
