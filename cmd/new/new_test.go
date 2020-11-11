package new

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommandFactory(t *testing.T) {
	c, err := CommandFactory()
	require.NoError(t, err)
	require.NoError(t, os.MkdirAll("testdata", 0755))
	require.NoError(t, os.Chdir("testdata"))
	i, err := c.Run(nil, nil)
	require.NoError(t, err)
	require.NotEqual(t, i, 0)
	_, err = c.Run([]string{"a/b"}, nil)
	require.NoError(t, err)
	require.NoError(t, os.Chdir("../"))
	require.NoError(t, os.RemoveAll("./b"))
}
