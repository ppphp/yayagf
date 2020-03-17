package _package

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/mitchellh/cli"
	"gitlab.papegames.com/fengche/yayagf/internal/command"
	"gitlab.papegames.com/fengche/yayagf/internal/file"
)

type Command struct {
}

func (c *Command) Help() string {
	return ""
}

func (c *Command) Synopsis() string {
	return "package a go project"
}

func (c *Command) Run(args []string) int {
	root, err := file.GetAppRoot()
	name := filepath.Base(root)
	if err != nil {
		log.Printf("get project name failed: %v", err.Error())
		return 1
	}

	priB, err := ioutil.ReadFile(os.Getenv("HOME") + "/.ssh/id_rsa")
	if err != nil {
		log.Println(err)
	}
	priS := string(priB)
	os.Setenv("pri", priS)

	pubB, err := ioutil.ReadFile(os.Getenv("HOME") + "/.ssh/id_rsa.pub")
	if err != nil {
		log.Println(err)
	}
	pubS := string(pubB)
	os.Setenv("pub", pubS)

	out, errs := &bytes.Buffer{}, &bytes.Buffer{}

	if err := command.DoCommand("docker", []string{"build", "-t", fmt.Sprintf("docker.papegames.com/%v", name), ".", "--build-arg", "pri", "--build-arg", "pub"}, out, errs); err != nil {
		log.Fatalf("docker build failed: %v", errs.String())
		return 1
	}
	if err := command.DoCommand("docker", []string{"save", fmt.Sprintf("docker.papegames.com/%v", name), "-o", fmt.Sprintf("%v.tar", name)}, out, errs); err != nil {
		log.Fatalf("docker save error: %v", errs.String())
		return 1
	}

	return 0
}

func CommandFactory() (cli.Command, error) {
	c := &Command{}
	return c, nil
}
