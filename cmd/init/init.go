package init

import (
	"os"

	"github.com/ppphp/yayagf/internal/file"
	"github.com/ppphp/yayagf/internal/log"
	"github.com/ppphp/yayagf/pkg/cli"
)

func CommandFactory() (*cli.Command, error) {
	return &cli.Command{Run: runInit}, nil
}

func runInit(args []string, flags map[string]string) (int, error) {
	if len(args) == 0 {
		log.Errorf("no project name")
		return 1, nil
	}
	mod, _, name := file.ExtractMod(args[0])
	dir, err := os.Getwd()
	if err != nil {
		log.Errorf("abs(%v) failed %v", name, err)
		return 1, err
	}

	return file.InitProject(mod, dir, name)
}
