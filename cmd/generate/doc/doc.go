package doc

import (
	"log"
	"os"

	"gitlab.papegames.com/fengche/yayagf/pkg/cli"
	"gitlab.papegames.com/fengche/yayagf/internal/command"
	"gitlab.papegames.com/fengche/yayagf/internal/file"
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

			if err, _, e := command.DoCommand2("swag", "init", "--output", "app/doc"); err != nil {
				log.Fatal(err, e)
			}

			return 0, nil
		},
	}, nil
}
