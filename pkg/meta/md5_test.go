package meta

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateSelfMD5(t *testing.T) {
	os.Args = []string{"sb"}
	_, err := CalculateSelfMD5()
	require.Error(t, err)
}
