package _package

import (
	"bytes"
	"fmt"
	"log"
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
	return "package"
}

func (c *Command) Run(args []string) int {
	root, err := file.GetAppRoot()
	name := filepath.Base(root)
	if err != nil {
		log.Printf("get project name failed: %v", err.Error())
		return 1
	}
	out, errs := &bytes.Buffer{}, &bytes.Buffer{}
	if err := command.DoCommand("docker", []string{"build", "-t", fmt.Sprintf("docker.papegames.com/%v", name), "."}, out, errs); err != nil {
		log.Fatalf("docker build failed: %v", errs.String())
		return 1
	}
	if err := command.DoCommand("docker", []string{"save", fmt.Sprintf("docker.papegames.com/%v", name), "-o", fmt.Sprintf("%v.tar", name)}, out, errs); err != nil {
		log.Fatalf("docker save error: %v", errs.String())
		return 1
	}

	return 0
}

func CommandFactory() (cli.Command, error) {
	c := &Command{}
	return c, nil
}
