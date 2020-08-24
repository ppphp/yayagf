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
var Meta = struct {
	// binary自带信息
	GoOS        string
	GoVersion   string
	GoArch      string
	GoCompiler  string
	GoBuildInfo *debug.BuildInfo

	MD5    string
	Uptime time.Time
	Mtime  time.Time
	Local  string
	// 注入
	BuildAt string
	Commit  string
	Version string
}{
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
