package generate

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/cli"
	"gitlab.papegames.com/fengche/yayagf/internal/util"
)

type Command struct {
}

func (c *Command) Help() string {
	return ""
}

func (c *Command) Synopsis() string {
	return "generator"
}

func (c *Command) Run(args []string) int {
	if len(args) == 0 {
		log.Println("need generate something")
		return 1
	}

	switch args[0] {
	case "table":
		pwd, err := os.Getwd()
		if err != nil {
			log.Panic(err)
		}
		root, err := util.FindAppRoot(pwd)
		if err != nil {
			log.Panic(err)
		}
		f, err := util.CreateFile(filepath.Join(root, "migrates"), false)
		if err != nil {
			log.Panic(err)
		}
		f.WriteString("")
	default:
		log.Println("need generate something")
		return 1
	}

	return 0
}

func CommandFactory() (cli.Command, error) {
	c := &Command{}
	return c, nil
}
