package server

import (
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCommandFactory(t *testing.T) {
	c, err := CommandFactory()
	require.NoError(t, err)
	require.NoError(t, os.Chdir("testdata/a"))
	go func() {
		i, err := c.Run(nil, nil)
		require.NoError(t, err)
		require.NotEqual(t, i, 0)
	}()
	time.Sleep(10 * time.Millisecond)

	_, err = os.Create("a.go")
	require.NoError(t, err)
}

func TestKill(t *testing.T) {
	cmd := exec.Command("sleep", "0.01")
	require.NoError(t, cmd.Start())
	require.NoError(t, kill(cmd))
}
