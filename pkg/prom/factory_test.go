package prom

import "testing"

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
}

func TestDBWaitCount(t *testing.T) {
	DBWaitCount("test", nil)
}

func TestDBWaitDuration(t *testing.T) {
	DBWaitDuration("test", nil)
}

func TestCallHTTPConnection(t *testing.T) {
	CallHTTPConnection()
}

func TestCallHTTPTTL(t *testing.T) {
	CallHTTPTTL()
}

func TestAll(t *testing.T) {

}
