package init

import (
	"path/filepath"

	"gitlab.papegames.com/fengche/yayagf/internal/file"
	"gitlab.papegames.com/fengche/yayagf/internal/log"
	"gitlab.papegames.com/fengche/yayagf/pkg/cli"
)

func CommandFactory() (*cli.Command, error) {
	return &cli.Command{Run: runInit}, nil
}

func runInit(args []string, flags map[string]string) (int, error) {
	if len(args) == 0 {
		log.Errorf("no project name")
		return 1, nil
	}
	namespace, name := filepath.Split(args[0])
	mod := filepath.Join(namespace, name)
	dir, err := filepath.Abs(filepath.Clean(name))
	if err != nil {
		log.Errorf("abs(%v) failed %v", name, err)
		return 1, err
	}

	return file.InitProject(mod, dir, name)
}
