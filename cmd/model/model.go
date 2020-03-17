package model

import (
	"bytes"
	"log"
	"os"
	"path/filepath"

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
	return "model issue for a yayagf project"
}

func (c *Command) Run(args []string) int {
	root, err := file.GetAppRoot()
	if err != nil {
		log.Printf("get project name failed: %v", err.Error())
		return 1
	}
	if err := os.Chdir(filepath.Join(root, "app")); err != nil {
		log.Printf("chdir failed: %v", err.Error())
		return 1
	}
	out, errs := &bytes.Buffer{}, &bytes.Buffer{}
	if err := command.DoCommand("entc", []string{"init"}, out, errs); err != nil {
		log.Fatalf("ent init failed: %v", errs.String())
		return 1
	}

	return 0
}

func CommandFactory() (cli.Command, error) {
	c := &Command{}
	return c, nil
}
