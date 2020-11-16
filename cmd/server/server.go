package server

import (
	"bytes"
	"fmt"
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
	return &cli.Command{Run: runServer}, nil
}

func runServer(args []string, flags map[string]string) (int, error) {
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
	for event := range c.watcher.Event {
		log.Printf("event: %v\n", event)

		fname, err := build.BuildBinary()
		if err != nil {
			log.Println(err.Error())
			continue
		}

		if c.cmd != nil && c.cmd.ProcessState != nil && !c.cmd.ProcessState.Exited() {
			if err := compareFile(fname, lastName); err != nil {
				continue
			}
			lastName = fname
		}
		if err := kill(c.cmd); err != nil {
			log.Printf("kill %v err: %v", c.cmd.Process.Pid, err)
		}
		c.cmd = command.GoCommand(fname, nil, os.Stdout, os.Stderr)
	}
	return 0, nil
}

func compareFile(fname, lastName string) error {
	f1, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	if lastName != "" {
		f2, err := ioutil.ReadFile(lastName)
		if err != nil {
			return err
		}
		if bytes.Equal(f1, f2) {
			return fmt.Errorf("equal")
		}
	}
	return nil
}

func kill(cmd *exec.Cmd) error {
	if cmd != nil && cmd.ProcessState == nil && cmd.Process != nil {
		if err := cmd.Process.Kill(); err != nil {
			return err
		}
	}
	return nil
}
