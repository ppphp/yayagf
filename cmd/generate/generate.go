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
	return "generate blueprint for a yayagf project"
}

func (c *Command) Run(args []string) int { return 0 }

func CommandFactory() (cli.Command, error) {
	c := &Command{}
	return c, nil
}
