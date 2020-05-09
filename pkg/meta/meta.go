// package for binary metadata like compilers and version
package meta

import (
	"runtime"
	"runtime/debug"
)

var BuildInfo, _ = debug.ReadBuildInfo()
var binhash, _ = CalculateSelfMD5()

// 一个构建期注入的meta信息结构体
var Meta = struct {
	// binary自带信息
	GoOS        string
	GoVersion   string
	GoArch      string
	GoCompiler  string
	GoBuildInfo *debug.BuildInfo

	MD5 string
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
}
