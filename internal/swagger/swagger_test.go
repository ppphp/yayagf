package swagger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSwagger(t *testing.T) {

	assert.NoError(t, os.Chdir("testdata"))
	assert.NoError(t, GenerateSwagger())
	assert.NoError(t, os.Chdir("../"))
	assert.Error(t, GenerateSwagger())
}
