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
	require.NoError(t, ioutil.WriteFile(filepath.Join(path, "go.mod"), []byte(`
module gitlab.papegames.com/fengche/yayagf

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/facebook/ent v0.1.2
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/lib/pq v1.3.0
	github.com/mitchellh/cli v1.0.0
	github.com/prometheus/client_golang v1.5.1
	github.com/sirupsen/logrus v1.4.2
	gitlab.papegames.com/fringe/quartz v0.0.0-20200103072440-229d00f9ada6
	golang.org/x/tools v0.0.0-20191012152004-8de300cfc20a
)

`), 0644))
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
