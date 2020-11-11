package version

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommandFactory(t *testing.T) {
	got, err := CommandFactory()
	require.NoError(t, err, "CommandFactory() error = %v", err)
	i, err := got.Run(nil, nil)
	require.NoError(t, err)
	require.Equal(t, i, 0)
}
