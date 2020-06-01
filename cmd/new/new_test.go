package new

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCommandFactory(t *testing.T) {
	c , err := CommandFactory()
	assert.NoError(t, err)
	assert.NoError(t, os.MkdirAll("testdata", 0755))
	assert.NoError(t, os.Chdir("testdata"))
	i, err := c.Run(nil, nil)
	assert.NoError(t, err)
	assert.NotEqual(t, i, 0)
	_, err = c.Run([]string{"a/b"}, nil)
	assert.NoError(t, err)
	assert.NoError(t, os.Chdir("../"))
	assert.NoError(t, os.RemoveAll("./b"))
}
