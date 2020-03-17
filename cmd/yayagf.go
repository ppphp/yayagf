package cmd

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
	"github.com/spf13/cobra"
	"gitlab.papegames.com/fengche/yayagf/cmd/doc"
	"gitlab.papegames.com/fengche/yayagf/cmd/generate"
	"gitlab.papegames.com/fengche/yayagf/cmd/interactive"
	"gitlab.papegames.com/fengche/yayagf/cmd/model"
	"gitlab.papegames.com/fengche/yayagf/cmd/new"
	_package "gitlab.papegames.com/fengche/yayagf/cmd/package"
	"gitlab.papegames.com/fengche/yayagf/cmd/server"
)

var RootCmd = &cobra.Command{
	Short: "yet another yet another go framework cli interface",
	Long:  `yet another yet another go framework cli interface just like rails`,
}

func Main() {
	c := cli.NewCLI("yayagf", "HEAD")

	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"doc":         doc.CommandFactory,
		"generate":    generate.CommandFactory,
		"interactive": interactive.CommandFactory,
		"model":       model.CommandFactory,
		"new":         new.CommandFactory,
		"package":     _package.CommandFactory,
		"server":      server.CommandFactory,
	}
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}
