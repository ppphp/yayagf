package prom

import (
	"net/http"
	"net/http/httptest"
	"runtime"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gomodule/redigo/redis"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/assert"
)

func TestSysCPU(t *testing.T) {
	SysCPU()
}

func TestSysMem(t *testing.T) {
	SysMem()
}

func TestSysDisk(t *testing.T) {
	SysDisk()
}

func TestSysLoad(t *testing.T) {
	SysLoad()
}

func TestGoRoutine(t *testing.T) {
	GoRoutine()
}

func TestGoMem(t *testing.T) {
	GoMem()
}

func TestGoGCTime(t *testing.T) {
	GoGCTime()
}

func TestRedisConnection(t *testing.T) {
	RedisConnection("test", nil)
}

func TestRedisWaitDuration(t *testing.T) {
	RedisWaitDuration("test", nil)
}

func TestRedisWaitCount(t *testing.T) {
	RedisWaitCount("test", nil)
}

func TestUrlTTL(t *testing.T) {
	URLTTL()
}

func TestUrlConnection(t *testing.T) {
	URLConnection()
}

func TestDbConnection(t *testing.T) {
	DbConnection("test", nil)
	DbConnection("root:root@(127.0.0.1:3306)/test?parseTime=true&loc=Local&charset=utf8", nil)
}

func TestDbClose(t *testing.T) {
	DbClose("test", nil)
	DbClose("root:root@(127.0.0.1:3306)/test?parseTime=true&loc=Local&charset=utf8", nil)
}

func TestDBWaitCount(t *testing.T) {
	DBWaitCount("test", nil)
	DBWaitCount("root:root@(127.0.0.1:3306)/test?parseTime=true&loc=Local&charset=utf8", nil)
}

func TestDBWaitDuration(t *testing.T) {
	DBWaitDuration("test", nil)
	DBWaitDuration("root:root@(127.0.0.1:3306)/test?parseTime=true&loc=Local&charset=utf8", nil)
}

func TestCallHTTPConnection(t *testing.T) {
	CallHTTPConnection()
}

func TestCallHTTPTTL(t *testing.T) {
	CallHTTPTTL()
}

func TestAll(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		c, _, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer c.Close()
		p := &redis.Pool{}
		runtime.GC()
		reg := prometheus.NewRegistry()

		reg.MustRegister(
			SysCPU(),
			SysMem(),
			SysDisk(),
			SysLoad(),
			GoRoutine(),
			GoMem(),
			GoGCTime(),
			RedisConnection("test", p),
			RedisWaitDuration("test", p),
			RedisWaitCount("test", p),
			// RedisConnection("e", nil),
			// RedisWaitDuration("e", nil),
			// RedisWaitCount("e", nil),
			URLTTL(),
			URLConnection(),
			DbConnection("test", c),
			DbClose("test", c),
			DBWaitCount("test", c),
			DBWaitDuration("test", c),
			// DbConnection("e", nil),
			// DbClose("e", nil),
			// DBWaitCount("e", nil),
			// DBWaitDuration("e", nil),
			CallHTTPConnection(),
			CallHTTPTTL(),
		)

		rr := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodGet, "http://0.0.0.0:8080/metrics", nil)
		assert.NoError(t, err)
		promhttp.InstrumentMetricHandler(
			reg, promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
		).ServeHTTP(rr, r)
		assert.NotEqual(t, "", rr.Result())
	})

	t.Run("nil", func(t *testing.T) {
		reg := prometheus.NewRegistry()

		reg.MustRegister(
			RedisConnection("e", nil),
			RedisWaitDuration("e", nil),
			RedisWaitCount("e", nil),
			DbConnection("e", nil),
			DbClose("e", nil),
			DBWaitCount("e", nil),
			DBWaitDuration("e", nil),
		)

		rr := httptest.NewRecorder()
		r, err := http.NewRequest(http.MethodGet, "http://0.0.0.0:8080/metrics", nil)
		assert.NoError(t, err)
		promhttp.InstrumentMetricHandler(
			reg, promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
		).ServeHTTP(rr, r)
		assert.NotEqual(t, "", rr.Result())
	})
}
