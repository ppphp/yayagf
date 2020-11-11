package cmd

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	got := Main()
	assert.Equal(t, got, 1)
}
