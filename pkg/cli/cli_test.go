package cli

import (
	"testing"
)

func TestApp(t *testing.T) {
	a := &App{"test", &Command{}}
	a.PrintMeta()
}

func TestCommand(t *testing.T){
	a := &App{"test", &Command{}}
	a.RawArgs = []string{"--version"}
	a.Run()
}
