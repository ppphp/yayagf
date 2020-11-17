package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadTomlFile(t *testing.T) {
	var k struct {
		Key int
	}
	require.NoError(t, LoadTomlFile("./testdata/conf.toml", &k))
	var k1 struct {
		Key string
	}
	require.Error(t, LoadTomlFile("./testdata/conf.toml", &k1))
}

func TestLoadEnv(t *testing.T) {
	var k struct {
		Key   int
		Value string
	}
	require.NoError(t, os.Setenv("KEY", "1"))
	require.NoError(t, os.Setenv("VALUE", "1"))
	LoadEnv(&k)

	func() {
		defer func() {
			require.NotNil(t, recover())
		}()

		var k struct {
			Key struct{}
		}

		LoadEnv(&k)

	}()
	require.NoError(t, os.Unsetenv("KEY"))
	require.NoError(t, os.Unsetenv("VALUE"))

}

func TestLoadConfig(t *testing.T) {
	var k struct {
		Key int
	}
	require.Error(t, LoadConfig(&k))

	require.NoError(t, os.Chdir("testdata"))

	require.NoError(t, LoadConfig(&k))
}
