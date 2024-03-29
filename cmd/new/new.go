package new

import (
	"os"
	"path/filepath"

	"github.com/ppphp/yayagf/internal/file"
	"github.com/ppphp/yayagf/internal/log"
	"github.com/ppphp/yayagf/pkg/cli"
)

func CommandFactory() (*cli.Command, error) {
	return &cli.Command{Run: runNew}, nil
}

func runNew(args []string, flags map[string]string) (int, error) {
	if len(args) == 0 {
		log.Errorf("no project name")
		return 1, nil
	}
	mod, _, name := file.ExtractMod(args[0])
	dir, err := filepath.Abs(filepath.Clean(name))
	if err != nil {
		log.Errorf("abs(%v) failed %v", name, err)
		return 1, err
	}
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Errorf("create (%v) failed %v", dir, err)
		return 1, err
	}
	if err := os.Chdir(dir); err != nil {
		log.Errorf("chdir (%v) failed %v", dir, err)
		return 1, err
	}

	return file.InitProject(mod, dir, name)
}
