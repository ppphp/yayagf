package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandFactory(t *testing.T) {
	got, err := CommandFactory()
	assert.NoError(t, err, "CommandFactory() error = %v", err)
	i, err := got.Run(nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, i, 0)
}
