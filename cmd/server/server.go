package server

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"

	"gitlab.papegames.com/fengche/yayagf/internal/command"
	"gitlab.papegames.com/fengche/yayagf/internal/file"
	"gitlab.papegames.com/fringe/quartz"
)

var Command = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		run(args)
	},
}

func run(args []string) int {
	type Command struct {
		watcher *quartz.Quartz
		pwd     string
		cmd     *exec.Cmd
	}
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("pwd error: %v", err.Error())
		return 1
	}
	watcher, err := quartz.NewQuartz(pwd, time.Second)
	if err != nil {
		return 1
	}
	c := &Command{watcher: watcher, pwd: pwd}
	root, err := file.GetAppRoot()
	if err != nil {
		log.Printf("GetAppRoot error: %v", err.Error())
		return 1
	}
	// specify build params
	os.Setenv("GOPROXY", "https://goproxy.io")
	os.Setenv("GOSUMDB", "off")
	os.Setenv("GOPRIVATE", "gitlab.papegames.com/*")

	// begin watch
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
			var o, e bytes.Buffer
			/*
				if err := command.DoCommand("swag", []string{"init", "-o", "app/docs"}, &o, &e); err != nil {
					log.Printf("swag to %v err: %v, err: %v, out: %v\n", "app/docs", e.String(), o.String())
					continue
				}
			*/
			if err := command.DoCommand("go", []string{"build", "-o", f.Name(), "./"}, &o, &e); err != nil {
				log.Printf("build to %v err: %v, err: %v, out: %v\n", f.Name(), err, e.String(), o.String())
				continue
			}
			if c.cmd != nil && c.cmd.ProcessState != nil && !c.cmd.ProcessState.Exited() {
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
			if c.cmd != nil && c.cmd.Process != nil {
				if err := c.cmd.Process.Kill(); err != nil {
					log.Printf("kill %v err: %v", c.cmd.Process.Pid, err)
				}
			}
			c.cmd = command.GoCommand(f.Name(), nil, os.Stdout, os.Stderr)
		}
	}

	return 0
}
