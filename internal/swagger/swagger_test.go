package swagger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateSwagger(t *testing.T) {

	require.NoError(t, os.Chdir("testdata"))
	require.NoError(t, GenerateSwagger())
	require.NoError(t, os.Chdir("../"))
	require.Error(t, GenerateSwagger())
}
