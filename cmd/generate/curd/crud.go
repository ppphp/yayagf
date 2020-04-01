package curd

import (
	"log"
	"path/filepath"

	"gitlab.papegames.com/fengche/yayagf/internal/ent"

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
			mod, err := file.GetMod(root)
			if err := ent.GenerateCRUDFiles(mod, filepath.Join(root, "app", "schema"), filepath.Join(root, "app", "crud")); err != nil {
				return 1, err
			}

			return 0, nil
		},
	}
	return c, nil
}
