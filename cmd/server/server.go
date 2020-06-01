package server

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"gitlab.papegames.com/fengche/yayagf/internal/command"
	"gitlab.papegames.com/fengche/yayagf/internal/file"
	"gitlab.papegames.com/fengche/yayagf/pkg/cli"
	"gitlab.papegames.com/fringe/quartz"
)

func CommandFactory() (*cli.Command, error) {
	c := &cli.Command{
		Run: func(args []string, flags map[string]string) (int, error) {
			type Command struct {
				watcher *quartz.Quartz
				pwd     string
				cmd     *exec.Cmd
			}
			pwd, err := os.Getwd()
			if err != nil {
				log.Printf("pwd error: %v", err.Error())
				return 1, err
			}
			watcher, err := quartz.NewQuartz(pwd, time.Second)
			if err != nil {
				return 1, err
			}
			c := &Command{watcher: watcher, pwd: pwd}
			root, err := file.GetAppRoot()
			if err != nil {
				log.Printf("GetAppRoot error: %v", err.Error())
				return 1, err
			}
			// specify build params
			_ = os.Setenv("GOPROXY", "https://goproxy.io")
			_ = os.Setenv("GOSUMDB", "off")
			_ = os.Setenv("GOPRIVATE", "gitlab.papegames.com/*")

			// begin watch
			_ = os.Chdir(root)
			if err := c.watcher.Begin(); err != nil {
				return 1, err
			}
			defer c.watcher.Stop()
			lastName := ""
			go func() { c.watcher.Event <- "" }()
			for {
				select {
				case event, ok := <-c.watcher.Event:
					if !ok {
						continue
					}
					log.Printf("event: %v\n", event)

					f, err := ioutil.TempFile("/tmp", "*")
					if err != nil {
					}
					f.Close()
					var o, e bytes.Buffer
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

		},
	}
	return c, nil
}
