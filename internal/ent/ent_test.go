package ent

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateSchema(t *testing.T) {
	require.NoError(t, os.Chdir("testdata"))
	require.NoError(t, GenerateSchema("./app/schema", []string{"b"}))
	require.NoError(t, os.Remove("./app/schema/b.go"))
	require.NoError(t, os.Chdir("../"))

}

func TestGenerateCRUDFiles(t *testing.T) {
	require.NoError(t, os.Chdir("testdata"))
	require.NoError(t, GenerateCRUDFiles("testdata", "./app/schema", "./app/crud", nil))
	require.NoError(t, os.Chdir("../"))
}
