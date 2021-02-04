package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetMod(t *testing.T) {
	path := os.TempDir()
	gomod, err := ioutil.ReadFile("../../go.mod")
	require.NoError(t, err)
	require.NoError(t, ioutil.WriteFile(filepath.Join(path, "go.mod"), gomod, 0644))
	mod, err := GetMod(path)
	require.NoError(t, err)

	require.Equal(t, "gitlab.papegames.com/fengche/yayagf", mod)
}

func TestGetAppRoot(t *testing.T) {
	root, err := GetAppRoot()
	require.NoError(t, err)
	require.NotEqual(t, "", root)

	require.NoError(t, os.RemoveAll("/tmp/go.mod"))
	require.NoError(t, os.Chdir("/tmp"))

	root, err = GetAppRoot()
	require.Error(t, err, "should error but %v", root)
	require.Equal(t, "", root)

}
