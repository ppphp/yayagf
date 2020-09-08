// 测量命名规则：
// 通用label: service_name
// 字段名称：按照功能等级划分，系统、运行时、存储、调用、服务
// 专用label：功能内部的平级监控
package prom

import (
	"database/sql"
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/gomodule/redigo/redis"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

func SysCPU() *GaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace:   "system",
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

func SysMem() *GaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace:   "system",
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

func SysDisk() *GaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace:   "system",
		Name:        "disk",
		ConstLabels: map[string]string{},
	}, []string{"path", "status"}, func() []LV {
		ps, err := disk.Partitions(false)
		if err != nil {
			return nil
		}
		lvs := []LV{}
		for _, v := range ps {
			d, err := disk.Usage(v.Mountpoint)
			if err == nil {
				lvs = append(lvs, LV{Lbs: []string{v.Mountpoint, "total"}, V: float64(d.Total)})
				lvs = append(lvs, LV{Lbs: []string{v.Mountpoint, "free"}, V: float64(d.Free)})
				lvs = append(lvs, LV{Lbs: []string{v.Mountpoint, "used"}, V: float64(d.Used)})
				lvs = append(lvs, LV{Lbs: []string{v.Mountpoint, "inode_used"}, V: float64(d.InodesUsed)})
				lvs = append(lvs, LV{Lbs: []string{v.Mountpoint, "inode_total"}, V: float64(d.InodesTotal)})
				lvs = append(lvs, LV{Lbs: []string{v.Mountpoint, "inode_free"}, V: float64(d.InodesFree)})
			}
		}
		return lvs
	})
}

func SysLoad() *GaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace:   "system",
		Name:        "load_average",
		ConstLabels: map[string]string{},
	}, []string{"cat"}, func() (lvs []LV) {
		lvs = []LV{}
		a, err := load.Avg()
		if err != nil {
			return
		}
		lvs = append(lvs, LV{Lbs: []string{"avg_1"}, V: a.Load1})
		lvs = append(lvs, LV{Lbs: []string{"avg_5"}, V: a.Load5})
		lvs = append(lvs, LV{Lbs: []string{"avg_15"}, V: a.Load15})
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
		Name:      "goroutines",
	}, func() float64 {
		return float64(runtime.NumGoroutine())
	})
}

// 注释 https://blog.csdn.net/m0_38132420/article/details/71699815
func GoMem() *GaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace: "runtime",
		Name:      "mem",
	}, []string{"cat"}, func() (lvs []LV) {
		lvs = []LV{}
		memstat := &runtime.MemStats{}
		runtime.ReadMemStats(memstat)
		lvs = append(lvs, LV{Lbs: []string{"sys"}, V: float64(memstat.Sys)})                //服务现在系统使用的内存
		lvs = append(lvs, LV{Lbs: []string{"alloc"}, V: float64(memstat.Alloc)})            //golang语言框架堆空间分配的字节数
		lvs = append(lvs, LV{Lbs: []string{"total_alloc"}, V: float64(memstat.TotalAlloc)}) //从服务开始运行至今分配器为分配的堆空间总和，只有增加，
		//释放的时候不减少
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

func GoGCTime() *GaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace: "runtime",
		Name:      "gc_time",
	}, []string{"quantile"}, func() []LV {
		gst := debug.GCStats{}
		debug.ReadGCStats(&gst)
		lvs := []LV{}
		for k, v := range gst.PauseQuantiles {
			lvs = append(lvs, LV{Lbs: []string{fmt.Sprint(k)}, V: v.Seconds()})
		}
		return lvs
	})
}

func URLTTL() *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   "service",
		Name:        "url_ttl",
		ConstLabels: map[string]string{},
	}, []string{"url", "method", "ret"})
}

func URLConnection() *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace:   "service",
		Name:        "http_connections",
		ConstLabels: map[string]string{},
	}, []string{"url", "method"})
}

func RedisConnection(dbname string, client *redis.Pool) *GaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace:   "redis",
		Name:        "redis_connections",
		ConstLabels: map[string]string{},
	}, []string{"dbname", "type"}, func() []LV {
		if client == nil {
			return nil
		}
		lvs := []LV{
			// The number of active connections.
			{[]string{dbname, "active"}, float64(client.Stats().ActiveCount)},
			// The number of open connections that are currently idle.
			{[]string{dbname, "idle"}, float64(client.Stats().IdleCount)},
		}
		return lvs
	})
}

func RedisWaitCount(dbname string, client *redis.Pool) *GaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace:   "redis",
		Name:        "wait_count",
		ConstLabels: map[string]string{},
	}, []string{"dbname"}, func() []LV {
		if client == nil {
			return nil
		}
		return []LV{{[]string{dbname}, float64(client.Stats().WaitCount)}}
	})
}

func RedisWaitDuration(dbname string, client *redis.Pool) *GaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace:   "redis",
		Name:        "wait_duration",
		ConstLabels: map[string]string{},
	}, []string{"dbname"}, func() []LV {
		if client == nil {
			return nil
		}
		return []LV{{[]string{dbname}, client.Stats().WaitDuration.Seconds()}}
	})
}

// https://orangematter.solarwinds.com/2018/05/22/new-stats-exposed-in-gos-database-sql-package/

// The number of connections.
func DbConnection(dbname string, client *sql.DB) *GaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace:   "db",
		Name:        "connection",
		ConstLabels: map[string]string{},
	}, []string{"dbname", "type"}, func() []LV {
		lvs := []LV{
			// The number of connections.
			{[]string{dbname, "open"}, float64(client.Stats().OpenConnections)},
			// The number of open connections that are currently idle.
			{[]string{dbname, "idle"}, float64(client.Stats().Idle)},
			// The number of connections actively in-use.
			{[]string{dbname, "inuse"}, float64(client.Stats().InUse)},
		}
		return lvs
	})
}

// The total number of times a goroutine has had to wait for a connection.
func DBWaitCount(dbname string, client *sql.DB) *GaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace:   "db",
		Name:        "wait_count",
		ConstLabels: map[string]string{},
	}, []string{"dbname"}, func() []LV {
		if client == nil {
			return nil
		}
		return []LV{{[]string{dbname}, float64(client.Stats().WaitCount)}}
	})
}

// The cumulative amount of time goroutines have spent waiting for a connection.
func DBWaitDuration(dbname string, client *sql.DB) *GaugeVecFuncCollector {
	return NewGaugeVecFunc(prometheus.GaugeOpts{
		Namespace:   "db",
		Name:        "wait_duration",
		ConstLabels: map[string]string{},
	}, []string{"dbname"}, func() []LV {
		if client == nil {
			return nil
		}
		return []LV{{[]string{dbname}, client.Stats().WaitDuration.Seconds()}}
	})
}

func CallHTTPConnection() *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace:   "call",
		Name:        "http_connections",
		ConstLabels: map[string]string{},
	}, []string{"service", "url", "method"})
}

func CallHTTPTTL() *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   "call",
		Name:        "http_ttl",
		ConstLabels: map[string]string{},
	}, []string{"service", "url", "method", "code", "ret"})
}
