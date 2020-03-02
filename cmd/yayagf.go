package cmd

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
	"gitlab.papegames.com/fengche/yayagf/cmd/generate"
	"gitlab.papegames.com/fengche/yayagf/cmd/interactive"
	"gitlab.papegames.com/fengche/yayagf/cmd/new"
	_package "gitlab.papegames.com/fengche/yayagf/cmd/package"
	"gitlab.papegames.com/fengche/yayagf/cmd/server"
)

func Main() {
	c := cli.NewCLI("yayagf", "HEAD")

	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"server":      server.CommandFactory,
		"new":         new.CommandFactory,
		"generate":    generate.CommandFactory,
		"package":     _package.CommandFactory,
		"interactive": interactive.CommandFactory,
	}
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}
