package cmd

import (
	"gitlab.papegames.com/fengche/yayagf/cmd/generate"
	"gitlab.papegames.com/fengche/yayagf/pkg/cli"
	"log"
	"os"

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
	}
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}
