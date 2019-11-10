package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
	"github.com/ppphp/yayagf/cmd/server"
)

func main() {
	c := cli.NewCLI("yayagf", "HEAD")

	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"server": server.CommandFactory,
	}
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}
