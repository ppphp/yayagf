// package for binary metadata like compilers and version
package meta

import (
	"net"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

var BuildInfo, _ = debug.ReadBuildInfo()
var binhash, _ = CalculateSelfMD5()
var uptime = time.Now()
var mtime = func() time.Time {
	st, err := os.Stat(os.Args[0])
	if err != nil {
		return time.Now()
	}
	return st.ModTime()
}()

var local = func() string {
	locals, err := net.Interfaces()
	if err != nil {
		panic("cannot find local address")
	}
	for _, interfac := range locals {
		if interfac.HardwareAddr.String() != "" {
			if strings.HasPrefix(interfac.Name, "en") ||
				strings.HasPrefix(interfac.Name, "eth") {
				if addrs, err := interfac.Addrs(); err == nil {
					for _, addr := range addrs {
						if addr.Network() == "ip+net" {
							pr := strings.Split(addr.String(), "/")
							if len(pr) == 2 && len(strings.Split(pr[0], ".")) == 4 {
								return pr[0]
							}
						}
					}
				}
			}
		}
	}
	panic("cannot find local address")
}()

// 一个构建期注入的meta信息结构体
type Meta struct {
	// binary自带信息

	GoOS        string // go运行系统
	GoVersion   string // go版本
	GoArch      string // go架构
	GoCompiler  string // go编译器
	GoBuildInfo *debug.BuildInfo

	MD5    string    // md5归纳
	Uptime time.Time // 启动时间
	Mtime  time.Time // 二进制修改时间
	Local  string    // 内网ip

	// 注入

	BuildAt string // 构建于
	Commit  string // 提交
	Version string // 大版本
}

func Get() Meta {
	return Meta{
		GoOS:        runtime.GOOS,
		GoVersion:   runtime.Version(),
		GoArch:      runtime.GOARCH,
		GoCompiler:  runtime.Compiler,
		GoBuildInfo: BuildInfo,
		Commit:      "HEAD",
		Version:     "HEAD",
		MD5:         binhash,
		Uptime:      uptime,
		Mtime:       mtime,
		Local:       local,
	}
}
