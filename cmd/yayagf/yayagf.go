package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
	"gitlab.papegames.com/fengche/yayagf/cmd/generate"
	"gitlab.papegames.com/fengche/yayagf/cmd/interactive"
	"gitlab.papegames.com/fengche/yayagf/cmd/new"
	"gitlab.papegames.com/fengche/yayagf/cmd/server"
)

func main() {
	c := cli.NewCLI("yayagf", "HEAD")

	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"server":      server.CommandFactory,
		"new":         new.CommandFactory,
		"generate":    generate.CommandFactory,
		"interactive": interactive.CommandFactory,
	}
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}
