package doc

import (
	"log"
	"os"

	"gitlab.papegames.com/fengche/yayagf/internal/swagger"

	"gitlab.papegames.com/fengche/yayagf/internal/file"
	"gitlab.papegames.com/fengche/yayagf/pkg/cli"
)

func CommandFactory() (*cli.Command, error) {
	return &cli.Command{
		Run: func(args []string) (int, error) {
			root, err := file.GetAppRoot()
			if err != nil {
				log.Fatal(err)
			}
			if err := os.Chdir(root); err != nil {
				log.Fatal(err)
			}

			if err := swagger.GenerateSwagger(); err != nil {
				log.Fatal(err)
			}

			return 0, nil
		},
	}, nil
}
