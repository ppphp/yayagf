package cli

import (
	"testing"

	"github.com/stretchr/testify/require"
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
