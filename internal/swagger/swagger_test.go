package swagger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateSwagger(t *testing.T) {

	wd, err := os.Getwd()
	require.NoError(t, err)
	require.NoError(t, os.Chdir("/tmp"))
	require.Error(t, GenerateSwagger())
	require.NoError(t, os.Chdir(wd))
	require.Error(t, GenerateSwagger(), "cannot generate at %v", wd)
	require.NoError(t, os.Chdir("testdata"))
	require.NoError(t, GenerateSwagger())
	require.NoError(t, os.Chdir("../"))
	require.Error(t, GenerateSwagger())
}
