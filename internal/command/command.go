package command

import (
	"io"
	"os/exec"
	"time"
)

func GoCommand(bin string, args []string, out io.Writer, err io.Writer) *exec.Cmd {
	cmd := exec.Command(bin, args...)
	cmd.Stdout = out
	cmd.Stderr = err

	go func() {
		cmd.Run()
	}()
	time.Sleep(time.Second)

	return cmd
}

func DoCommand(bin string, args []string, out io.Writer, err io.Writer) error {
	cmd := exec.Command(bin, args...)
	cmd.Stdout = out
	cmd.Stderr = err

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
