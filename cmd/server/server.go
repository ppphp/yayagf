package server

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/mitchellh/cli"
	"gitlab.papegames.com/fengche/quartz"
	"gitlab.papegames.com/fengche/yayagf/internal/command"
	"gitlab.papegames.com/fengche/yayagf/internal/util"
)

type Command struct {
	watcher *quartz.Quartz
	pwd     string
	pcs     *os.Process
}

func (c *Command) Help() string {
	return ""
}

func (c *Command) Synopsis() string {
	return "monitor your change, rebuild and run app"
}

func (c *Command) Run(args []string) int {
	wd, _ := os.Getwd()
	root, _ := util.FindAppRoot(wd)
	os.Chdir(root)
	c.watcher.Begin()
	defer c.watcher.Stop()
	lastName := ""
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
			cmd := exec.Command("go", "build", "-o", f.Name(), "./")
			var o, e bytes.Buffer
			cmd.Stdout = &o
			cmd.Stderr = &e
			if err := cmd.Run(); err != nil {
				log.Printf("build to %v err: %v, err: %v, out: %v\n", f.Name(), err, e.String(), o.String())
				continue
			}
			if !c.cmd.ProcessState.Exited() {
				f1, err1 := ioutil.ReadFile(f.Name())
				if err1 != nil {
					continue
				}
				if lastName != "" {
					f2, err2 := ioutil.ReadFile(lastName)
					if err2 != nil {
						continue
					}
					if bytes.Equal(f1, f2) {
						continue
					}
				}
				lastName = f.Name()
			}
			if c.pcs != nil {
				if err := c.pcs.Kill(); err != nil {
					log.Printf("kill %v err: %v", c.pcs.Pid, err)
				}
			}
			c.pcs = command.GoCommand(f.Name(), nil, os.Stdout, os.Stderr)
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
