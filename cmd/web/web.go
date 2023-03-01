package web

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/ppphp/yayagf/internal/blueprint"
	"github.com/ppphp/yayagf/internal/command"
	"github.com/ppphp/yayagf/internal/file"
	"github.com/ppphp/yayagf/internal/log"
	"github.com/ppphp/yayagf/pkg/cli"
)

func CommandFactory() (*cli.Command, error) {
	return &cli.Command{Run: runWeb}, nil
}

func runWeb(args []string, flags map[string]string) (int, error) {
	root, err := file.GetAppRoot()
	if err != nil {
		log.Errorf("no project")
		return 1, nil
	}
	_ = os.Setenv("GOOS", "js")
	defer os.Unsetenv("GOOS")
	_ = os.Setenv("GOARCH", "wasm")
	defer os.Unsetenv("GOARCH")
	err, content, se := command.DoCommand("go", "build", "-ldflags=-s -w", "-o" /*"./web/app.wasm") */, "/dev/stdout")
	if err != nil {
		log.Errorf("build error err (%v)", se)
		return 1, err
	}

	// useless but can be here
	_ = blueprint.WriteFileWithTmpl(filepath.Join(root, "app", "wasm", "wasm.go"), `package wasm

const WASM = {{.Content}}
`, struct{ Content string }{strconv.Quote(content)})

	return 0, nil
}
