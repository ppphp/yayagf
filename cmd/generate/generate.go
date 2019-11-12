package generate

import (
	"github.com/mitchellh/cli"
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
		println("need generate something")
		return 1
	}

	return 0
}

func CommandFactory() (cli.Command, error) {
	c := &Command{}
	return c, nil
}
