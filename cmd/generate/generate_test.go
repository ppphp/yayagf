package generate

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCommandFactory(t *testing.T) {
	_, err := CommandFactory()
	require.NoError(t, err)
}
