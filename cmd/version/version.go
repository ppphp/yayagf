package version

import (
	"github.com/ppphp/yayagf/pkg/cli"
)

func CommandFactory() (*cli.Command, error) {
	c := &cli.Command{
		Run: func(args []string, flags map[string]string) (int, error) {
			return 0, nil
		},
	}

	return c, nil
}
