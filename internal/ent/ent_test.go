package ent

import (
	"os"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGenerateSchema(t *testing.T) {
	assert.NoError(t, os.Chdir("testdata"))
	assert.NoError(t, GenerateSchema("./app/schema", []string{"b"}))
	assert.NoError(t, os.Remove("./app/schema/b.go"))
	assert.NoError(t, os.Chdir("../"))

}

func TestGenerateCRUDFiles(t *testing.T) {
	assert.NoError(t, os.Chdir("testdata"))
	assert.NoError(t, GenerateCRUDFiles("testdata", "./app/schema", "./app/crud", nil))
	assert.NoError(t, os.Chdir("../"))
}
