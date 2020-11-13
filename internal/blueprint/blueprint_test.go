package blueprint

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteFileWithTmpl(t *testing.T) {
	require.NoError(t, WriteFileWithTmpl("testdata/a", "abc", nil))
}
