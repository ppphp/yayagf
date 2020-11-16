package command

import (
	"bytes"
	"io"
	"os/exec"
)

func GoCommand(bin string, args []string, out io.Writer, err io.Writer) *exec.Cmd {
	cmd := exec.Command(bin, args...)
	cmd.Stdout = out
	cmd.Stderr = err

	start := make(chan struct{})
	go func() {
		close(start)
		_ = cmd.Run()
	}()
	<-start

	return cmd
}

func DoCommand(bin string, args ...string) (error, string, string) {
	out, errs := &bytes.Buffer{}, &bytes.Buffer{}
	cmd := exec.Command(bin, args...)
	cmd.Stdout = out
	cmd.Stderr = errs
	if err := cmd.Run(); err != nil {
		return err, out.String(), errs.String()
	}
	return nil, out.String(), errs.String()
}
