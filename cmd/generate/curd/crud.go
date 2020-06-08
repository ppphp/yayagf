package curd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"gitlab.papegames.com/fengche/yayagf/internal/log"
	"path/filepath"

	"gitlab.papegames.com/fengche/yayagf/internal/ent"

	"gitlab.papegames.com/fengche/yayagf/internal/file"
	"gitlab.papegames.com/fengche/yayagf/pkg/cli"
)

func CommandFactory() (*cli.Command, error) {
	c := &cli.Command{
		Run: func(args []string, flags map[string]string) (int, error) {
			templates := []string{}
			debug := false
			pf := pflag.NewFlagSet("", pflag.PanicOnError)
			pf.BoolVarP(&debug, "debug", "d", false, "")
			pf.StringArrayVarP(&templates, "template", "t", nil, "")
			if err := pf.Parse(args); err != nil {
				panic(err)
			}
			log.Debugf("%v", templates)
			if debug {
				log.Logger.SetReportCaller(true)
				log.Logger.SetLevel(logrus.DebugLevel)
			}

			log.Debugf("%v", args)
			root, err := file.GetAppRoot()
			if err != nil {
				log.Printf("get project name failed: %v", err.Error())
				return 1, err
			}
			mod, err := file.GetMod(root)
			if err != nil {
				return 1, err
			}
			if err := ent.GenerateCRUDFiles(mod, filepath.Join(root, "app", "schema"), filepath.Join(root, "app", "crud"), templates); err != nil {
				return 1, err
			}

			return 0, nil
		},
	}
	return c, nil
}
