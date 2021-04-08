// The composer of subcommands, using gitlab.papegames/fengche/yayagf/pkg/cli
package cmd

import (
	"os"

	_init "gitlab.papegames.com/fengche/yayagf/cmd/init"
	"gitlab.papegames.com/fengche/yayagf/cmd/version"
	"gitlab.papegames.com/fengche/yayagf/cmd/web"

	"gitlab.papegames.com/fengche/yayagf/cmd/generate"
	"gitlab.papegames.com/fengche/yayagf/pkg/cli"

	"gitlab.papegames.com/fengche/yayagf/cmd/new"
	_package "gitlab.papegames.com/fengche/yayagf/cmd/package"
	"gitlab.papegames.com/fengche/yayagf/cmd/server"
	"gitlab.papegames.com/fengche/yayagf/internal/log"
)

func Main() int {
	c := &cli.App{Name: "yayagf", Command: &cli.Command{}}

	if len(os.Args) > 0 {
		c.RawArgs = os.Args[1:]
	}

	c.Commands = map[string]cli.CommandFactory{
		"generate": generate.CommandFactory, "g": generate.CommandFactory,
		//"interactive": interactive.CommandFactory,
		"new":     new.CommandFactory,
		"web":     web.CommandFactory,
		"init":    _init.CommandFactory,
		"package": _package.CommandFactory,
		"server":  server.CommandFactory, "s": server.CommandFactory,
		"version": version.CommandFactory,
	}
	exitStatus, err := c.Run()

	if err != nil {
		log.Println(err)
	}

	return exitStatus
}
