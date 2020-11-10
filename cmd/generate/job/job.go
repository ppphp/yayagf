package job

import (
	"log"
	"os"
	"path/filepath"

	"gitlab.papegames.com/fengche/yayagf/internal/file"
	"gitlab.papegames.com/fengche/yayagf/pkg/cli"
)

func CommandFactory() (*cli.Command, error) {
	c := &cli.Command{
		Run: func(args []string, flags map[string]string) (int, error) {
			root, err := file.GetAppRoot()
			if err != nil {
				log.Printf("get project name failed: %v", err.Error())
				return 1, err
			}
			if len(args) == 0 {
				return 0, nil
			}
			if err := os.Mkdir(filepath.Join(root, "app", "jobs"), 0755); err != nil {
				return 0, err
			}
			for _, n := range args {
				f, err := os.OpenFile(filepath.Join(root, "app", "jobs", n), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
				if err != nil {
					return 1, err
				}
				if _, err := f.Write([]byte(`package jobs`)); err != nil {
					return 1, err
				}
			}

			return 0, nil
		},
	}
	return c, nil
}
