package doc

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
	"gitlab.papegames.com/fengche/yayagf/internal/command"
	"gitlab.papegames.com/fengche/yayagf/internal/file"
)

type Command struct {
}

func (c *Command) Help() string {
	return ""
}

func (c *Command) Synopsis() string {
	return "doc issue for a go project"
}

func (c *Command) Run(args []string) int {
	root, err := file.GetAppRoot()
	if err != nil {
		log.Fatal(err)
	}
	if err := os.Chdir(root); err != nil {
		log.Fatal(err)
	}

	if err, _, e := command.DoCommand2("swag", "init", "--output", "app/doc"); err != nil {
		log.Fatal(err, e)
	}

	return 0
}

func CommandFactory() (cli.Command, error) {
	c := &Command{}
	return c, nil
}
