package prom

import "testing"

func TestSysCPU(t *testing.T) {
	SysCPU("test")
}

func TestSysMem(t *testing.T) {
	SysMem("test")
}

func TestSysDisk(t *testing.T) {
	SysDisk("test")
}

func TestSysLoad(t *testing.T) {
	SysLoad("test")
}

func TestGoRoutine(t *testing.T) {
	GoRoutine("test")
}

func TestGoMem(t *testing.T) {
	GoMem("test")
}

func TestGoGCTime(t *testing.T) {
	GoGCTime("test")
}

func TestRedisConnection(t *testing.T) {
	RedisConnection("test", "test", nil)
}

func TestRedisWaitDuration(t *testing.T) {
	RedisWaitDuration("test", "test", nil)
}

func TestRedisWaitCount(t *testing.T) {
	RedisWaitCount("test", "test", nil)
}

func TestUrlTTL(t *testing.T) {
	UrlTTL("test")
}

func TestUrlConnection(t *testing.T) {
	UrlConnection("test")
}

func TestDbConnection(t *testing.T) {
	DbConnection("test", "test", nil)
}

func TestDBWaitCount(t *testing.T) {
	DBWaitCount("test", "test", nil)
}

func TestDBWaitDuration(t *testing.T) {
	DBWaitDuration("test", "test", nil)
}

func TestCallHTTPConnection(t *testing.T) {
	CallHTTPConnection("test", "test")
}

func TestCallHTTPTTL(t *testing.T) {
	CallHTTPTTL("test", "test")
}
