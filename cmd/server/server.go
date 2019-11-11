package server

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/mitchellh/cli"
	"github.com/ppphp/quartz"
)

type Command struct {
	watcher *quartz.Quartz
	pwd     string
}

func (c *Command) Help() string {
	return ""
}

func (c *Command) Synopsis() string {
	return ""
}

func (c *Command) Run(args []string) int {
	c.watcher.Begin()
	defer c.watcher.Stop()
	for {
		select {
		case event, ok := <-c.watcher.Event:
			if !ok {
				continue
			}
			log.Printf("event: %v\n", event)

			/*
				cmd := exec.Command("go", "test", "./...")
				if err := cmd.Run(); err != nil {
					log.Printf("test err: %v \n", err)
					continue
				}
			*/
			f, err := ioutil.TempFile("/tmp", "*")
			if err != nil {
			}
			f.Close()
			cmd := exec.Command("go", "build", "-o", f.Name(), "./cmd/...")
			if err := cmd.Run(); err != nil {
				log.Printf("build to %v err: %v\n", f.Name(), err)
				continue
			}
		}
	}

	return 0
}

func CommandFactory() (cli.Command, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	watcher, err := quartz.NewQuartz(pwd, time.Second)
	if err != nil {
		return nil, err
	}
	c := &Command{watcher: watcher, pwd: pwd}
	return c, nil
}
