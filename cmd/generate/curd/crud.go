package curd

import (
	"bytes"
	"log"
	"os"
	"path/filepath"

	"gitlab.papegames.com/fengche/yayagf/pkg/cli"
	"gitlab.papegames.com/fengche/yayagf/internal/command"
	"gitlab.papegames.com/fengche/yayagf/internal/file"
)

func CommandFactory() (*cli.Command, error) {
	c := &cli.Command{
		Run: func(args []string) (int, error) {
			root, err := file.GetAppRoot()
			if err != nil {
				log.Printf("get project name failed: %v", err.Error())
				return 1, err
			}
			if err := os.Chdir(filepath.Join(root, "app")); err != nil {
				log.Printf("chdir failed: %v", err.Error())
				return 1, err
			}

			out, errs := &bytes.Buffer{}, &bytes.Buffer{}
			if err := command.DoCommand("entc", []string{"generate", "./ent/schema"}, out, errs); err != nil {
				log.Fatalf("ent generate failed: %v", errs.String())
				return 1, err
			}

			return 0, nil
		},
	}
	return c, nil
}
