package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	got := Main()
	assert.Equal(t, got, 1)
}
