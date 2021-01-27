package curd

import (
	"github.com/sirupsen/logrus"
	"gitlab.papegames.com/fengche/yayagf/internal/log"
	"os"
	"path/filepath"

	"gitlab.papegames.com/fengche/yayagf/internal/ent"

	"gitlab.papegames.com/fengche/yayagf/internal/file"
	"gitlab.papegames.com/fengche/yayagf/pkg/cli"
)

func CommandFactory() (*cli.Command, error) {
	c := &cli.Command{
		Run: func(args []string, flags map[string]string) (int, error) {
			template := flags["t"]
			if template == "" {
				template = flags["template"]
			}
			templates := []string{}
			if template != "" {
				templates = append(templates, template)
			}
			_, debug := flags["d"]
			if debug == false {
				_, debug = flags["debug"]
			}
			if debug {
				log.Logger.SetReportCaller(true)
				log.Logger.SetLevel(logrus.DebugLevel)
			}

			if debug {
				log.Printf("%v", args)
			}
			root, err := file.GetAppRoot()
			if err != nil {
				log.Printf("get project name failed: %v", err.Error())
				return 1, err
			}
			mod, err := file.GetMod(root)
			if err != nil {
				return 1, err
			}
			if st, err := os.Stat(filepath.Join(root, "app", "schema", "template")); err == nil {
				if st.IsDir() {
					templates = append(templates, filepath.Join(root, "app", "schema", "template"))
				}
			}
			if debug {
				log.Printf("%v", templates)
			}
			if err := ent.GenerateCRUDFiles(mod, filepath.Join(root, "app", "schema"), filepath.Join(root, "app", "crud"), templates); err != nil {
				return 1, err
			}

			return 0, nil
		},
	}
	return c, nil
}
