package schema

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/ppphp/yayagf/internal/ent"

	"github.com/ppphp/yayagf/internal/file"
	"github.com/ppphp/yayagf/pkg/cli"
)

func CommandFactory() (*cli.Command, error) {
	c := &cli.Command{
		Run: func(args []string, flags map[string]string) (int, error) {
			root, err := file.GetAppRoot()
			if err != nil {
				log.Printf("get project name failed: %v", err.Error())
				return 1, err
			}
			schemas := []string{}
			for _, a := range args {
				schemas = append(schemas, strings.Title(a))
			}
			if err := ent.GenerateSchema(filepath.Join(root, "app", "schema"), schemas); err != nil {
				return 1, err
			}

			return 0, nil
		},
	}
	return c, nil
}
