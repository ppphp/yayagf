package ci

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCommandFactory(t *testing.T) {
	c, err := CommandFactory()
	if err != nil {
		assert.NoError(t, err)
	}
	if err := os.Chdir("./testdata/a"); err != nil {
		assert.NoError(t, err)
	}
	i, err := c.Run(nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, i, 0)
	st, err := os.Stat("Jenkinsfile")
	assert.NoError(t, err)
	assert.NotNil(t, st)
}
