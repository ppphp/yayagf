package server

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"gitlab.papegames.com/fengche/yayagf/internal/build"
	"gitlab.papegames.com/fengche/yayagf/internal/command"
	"gitlab.papegames.com/fengche/yayagf/internal/file"
	"gitlab.papegames.com/fengche/yayagf/pkg/cli"
	"gitlab.papegames.com/fengche/yayagf/pkg/watcher"
)

func CommandFactory() (*cli.Command, error) {
	c := &cli.Command{
		Run: func(args []string, flags map[string]string) (int, error) {
			type Command struct {
				watcher *watcher.Watcher
				pwd     string
				cmd     *exec.Cmd
			}
			pwd, err := os.Getwd()
			if err != nil {
				log.Printf("pwd error: %v", err.Error())
				return 1, err
			}
			watch, err := watcher.NewWatcher(pwd, time.Second)
			if err != nil {
				log.Fatal(err)
			}
			c := &Command{watcher: watch, pwd: pwd}
			root, err := file.GetAppRoot()
			if err != nil {
				log.Printf("GetAppRoot error: %v", err.Error())
				return 1, err
			}

			// begin watch
			_ = os.Chdir(root)
			if err := c.watcher.Begin(); err != nil {
				return 1, err
			}
			defer c.watcher.Stop()
			lastName := ""
			go func() { c.watcher.Event <- watcher.Event{} }()
			for {
				select {
				case event, ok := <-c.watcher.Event:
					if !ok {
						continue
					}
					log.Printf("event: %v\n", event)

					f, err := build.BuildBinary()
					if err != nil {
						log.Println(err.Error())
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

		},
	}
	return c, nil
}
