package ci

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommandFactory(t *testing.T) {
	c, err := CommandFactory()
	if err != nil {
		require.NoError(t, err)
	}
	if err := os.Chdir("./testdata/a"); err != nil {
		require.NoError(t, err)
	}
	i, err := c.Run(nil, nil)
	require.NoError(t, err)
	require.Equal(t, i, 0)
	st, err := os.Stat("Jenkinsfile")
	require.NoError(t, err)
	require.NotNil(t, st)
}
