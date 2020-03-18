package curd

import (
	"gitlab.papegames.com/fengche/yayagf/internal/ent"
	"log"
	"path/filepath"

	"gitlab.papegames.com/fengche/yayagf/internal/file"
	"gitlab.papegames.com/fengche/yayagf/pkg/cli"
)

func CommandFactory() (*cli.Command, error) {
	c := &cli.Command{
		Run: func(args []string) (int, error) {
			root, err := file.GetAppRoot()
			if err != nil {
				log.Printf("get project name failed: %v", err.Error())
				return 1, err
			}
			if err := ent.GenerateCRUDFiles(filepath.Join(root, "app", "schema"), filepath.Join(root, "app", "crud")); err != nil {
				return 1, err
			}

			return 0, nil
		},
	}
	return c, nil
}
