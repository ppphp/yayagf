package schema

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCommandFactory(t *testing.T) {
	c, err := CommandFactory()
	assert.NoError(t, err)
	assert.NoError(t, os.Chdir("testdata/a"))
	_, err = c.Run([]string{"a"}, nil)
	assert.NoError(t, err)
	assert.NoError(t, os.RemoveAll("app/schema/a.go"))
}
