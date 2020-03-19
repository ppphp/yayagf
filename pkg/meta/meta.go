package meta

import (
	"runtime"
	"runtime/debug"
)

var BuildInfo, _ = debug.ReadBuildInfo()

// 一个构建期注入的meta信息结构体
var Meta = struct {
	// binary自带信息
	GoOS        string
	GoVersion   string
	GoArch      string
	GoCompiler  string
	GoBuildInfo *debug.BuildInfo
	CPUProcs    int

	// 注入
	Commit  string
	Version string
}{
	GoOS:        runtime.GOOS,
	GoVersion:   runtime.Version(),
	GoArch:      runtime.GOARCH,
	GoCompiler:  runtime.Compiler,
	GoBuildInfo: BuildInfo,
	CPUProcs:    runtime.NumCPU(),
	Commit:      "HEAD",
	Version:     "HEAD",
}
