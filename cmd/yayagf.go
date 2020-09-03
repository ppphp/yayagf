package cmd

import (
	"log"
	"os"

	"gitlab.papegames.com/fengche/yayagf/cmd/version"

	"gitlab.papegames.com/fengche/yayagf/cmd/generate"
	"gitlab.papegames.com/fengche/yayagf/pkg/cli"

	"gitlab.papegames.com/fengche/yayagf/cmd/new"
	_package "gitlab.papegames.com/fengche/yayagf/cmd/package"
	"gitlab.papegames.com/fengche/yayagf/cmd/server"
)

func Main() {
	c := cli.NewApp("yayagf", "HEAD")

	if len(os.Args) > 0 {
		c.Args = os.Args[1:]
	}

	c.Commands = map[string]cli.CommandFactory{
		"generate": generate.CommandFactory, "g": generate.CommandFactory,
		//"interactive": interactive.CommandFactory,
		"new":     new.CommandFactory,
		"package": _package.CommandFactory,
		"server":  server.CommandFactory,
		"version": version.CommandFactory,
	}
	exitStatus, err := c.Run()

	if err != nil {
		log.Println(err)
	}

	if exitStatus != 0 {
		os.Exit(exitStatus)
	}
}
