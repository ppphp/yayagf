package swagger

import (
	"os"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestGenerateSwagger(t *testing.T) {

	assert.NoError(t, os.Chdir("testdata"))
	assert.NoError(t, GenerateSwagger())
	assert.NoError(t, os.Chdir("../"))
	assert.Error(t, GenerateSwagger())
}
