package cmd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	got := Main()
	require.Equal(t, got, 1)
}
