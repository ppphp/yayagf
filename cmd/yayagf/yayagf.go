package main

import "github.com/mitchelleh/cli"

func main() {
	c := cli.NewCLI("yayagf", "HEAD")
	
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
	}
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}
