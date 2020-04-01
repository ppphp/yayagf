package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGetMod(t *testing.T) {
	path := os.TempDir()
	ioutil.WriteFile(filepath.Join(path, "go.mod"), []byte(`
module gitlab.papegames.com/fengche/yayagf

go 1.13

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/facebookincubator/ent v0.1.2
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/lib/pq v1.3.0
	github.com/mitchellh/cli v1.0.0
	github.com/prometheus/client_golang v1.5.1
	github.com/sirupsen/logrus v1.4.2
	gitlab.papegames.com/fringe/quartz v0.0.0-20200103072440-229d00f9ada6
	golang.org/x/tools v0.0.0-20191012152004-8de300cfc20a
)

`), 0644)
	mod, err := GetMod(path)
	if err != nil {
		t.Errorf("%v", err)
	} else if mod != "gitlab.papegames.com/fengche/yayagf" {
		t.Errorf("%v (evaluated) != %v (expected)", mod, "gitlab.papegames.com/fengche/yayagf")
	}
}
