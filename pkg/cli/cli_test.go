package cli

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestApp(t *testing.T) {
	a := &App{"test", &Command{}}
	a.PrintMeta()
}

func TestCommand(t *testing.T) {
	a := &App{"test", &Command{}}
	a.RawArgs = []string{"--version"}
	_, err := a.Run()
	require.NoError(t, err)
}
