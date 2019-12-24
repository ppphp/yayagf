package command

import (
	"io"
	"os"
	"os/exec"
	"time"
)

func GoCommand(bin string, args []string, out io.Writer, err io.Writer) *os.Process {
	cmd := exec.Command(bin, args...)
	cmd.Stdout = out
	cmd.Stderr = err

	go func() {
		cmd.Run()
	}()
	time.Sleep(time.Second)

	return cmd.Process
}
