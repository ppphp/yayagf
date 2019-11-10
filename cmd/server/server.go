package server

import (
	"context"
	"log"
	"os"
	"os/exec"

	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/cli"
)

type Command struct {
	watcher *fsnotify.Watcher
	pwd     string
	ctx     context.Context
	cancel  func()
}

func (c *Command) Help() string {
	return ""
}

func (c *Command) Synopsis() string {
	return ""
}

func (c *Command) Run(args []string) int {
	defer c.watcher.Close()
	for {
		select {
		case event, ok := <-c.watcher.Events:
			if !ok {
				continue
			}
			log.Printf("event: %v\n", event)

			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("modified file:", event.Name)
			}
			cmd := exec.Cmd{
				Path: c.pwd,
				Args: []string{"go", "test", "./..."},
				Env:  os.Environ(),
			}
			if err := cmd.Run(); err != nil {
				log.Printf("err: %v\n", err)
				continue
			}
			if c.cancel != nil {
				c.cancel()
			}
			ctx, f := context.WithCancel(c.ctx)
			go func(ctx context.Context) {
				cmd := exec.Cmd{
					Path: c.pwd,
					Args: []string{"go", "build", "./cmd/..."},
					Env:  os.Environ(),
				}
				if err := cmd.Run(); err != nil {
					if c.cancel != nil {
						log.Printf("err: %v\n", err)
						c.cancel()
					}
				}
			}(ctx)
			c.cancel = f

		case err, ok := <-c.watcher.Errors:
			if !ok {
				continue
			}
			log.Printf("error: %v\n", err)
		}
	}

	return 0
}

func CommandFactory() (cli.Command, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	watcher.Add(pwd)
	c := &Command{watcher: watcher, pwd: pwd, ctx: context.Background()}
	return c, nil
}
