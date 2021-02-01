package curd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommandFactory(t *testing.T) {
	c, err := CommandFactory()
	require.NoError(t, err)
	require.NoError(t, os.Chdir("testdata/a"))
	_, err = c.Run([]string{"--debug"}, nil)
	require.NoError(t, err)
	require.NoError(t, os.RemoveAll("app/crud"))
}
