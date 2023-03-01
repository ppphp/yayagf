// The composer of subcommands, using gitlab.papegames/fengche/yayagf/pkg/cli
package cmd

import (
	"os"

	_init "github.com/ppphp/yayagf/cmd/init"
	"github.com/ppphp/yayagf/cmd/version"
	"github.com/ppphp/yayagf/cmd/web"

	"github.com/ppphp/yayagf/cmd/generate"
	"github.com/ppphp/yayagf/pkg/cli"

	"github.com/ppphp/yayagf/cmd/new"
	_package "github.com/ppphp/yayagf/cmd/package"
	"github.com/ppphp/yayagf/cmd/server"
	"github.com/ppphp/yayagf/internal/log"
)

func Main() int {
	c := &cli.App{Name: "yayagf", Command: &cli.Command{}}

	if len(os.Args) > 0 {
		c.RawArgs = os.Args[1:]
	}

	c.Commands = map[string]cli.CommandFactory{
		"generate": generate.CommandFactory, "g": generate.CommandFactory,
		//"interactive": interactive.CommandFactory,
		"new":  new.CommandFactory,
		"web":  web.CommandFactory,
		"init": _init.CommandFactory, "i": _init.CommandFactory,
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
