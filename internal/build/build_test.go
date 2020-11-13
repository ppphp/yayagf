package build

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildBinary(t *testing.T) {
	require.NoError(t, os.Chdir("../../"))
	fname, err := BuildBinary()
	require.NoError(t, err, "%v", err)
	require.NotNil(t, fname)
}
