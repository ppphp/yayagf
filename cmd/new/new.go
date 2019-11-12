package new

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mitchellh/cli"
)

type Command struct {
}

func (c *Command) Help() string {
	return ""
}

func (c *Command) Synopsis() string {
	return "init a yayagf project"
}

func (c *Command) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println("no project name")
		return 1
	}
	dir, file := filepath.Split(args[0])
	name, err := filepath.Abs(filepath.Clean(file))
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	if err := os.Mkdir(name, 0744); err != nil {
		fmt.Println(err.Error())
		return 1
	}
	if err := os.Chdir(name); err != nil {
		fmt.Println(err.Error())
		return 1
	}

	cmd := exec.Command("go", "mod", "init", filepath.Join(dir, file))
	cmd.Run()

	if err := os.Mkdir(filepath.Join(name, "app"), 0744); err != nil {
		fmt.Println(err.Error())
		return 1
	}

	if f, err := os.OpenFile("main.go", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644); err != nil {
		fmt.Println(err.Error())
		return 1
	} else {
		f.Write([]byte(`package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello, world!\n")
}


		`))
	}

	return 0
}

func CommandFactory() (cli.Command, error) {
	c := &Command{}
	return c, nil
}
