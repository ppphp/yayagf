package doc

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"gitlab.papegames.com/fengche/yayagf/internal/command"
	"gitlab.papegames.com/fengche/yayagf/internal/file"
)

var Command = &cobra.Command{
	Use: "doc",
	Run: func(cmd *cobra.Command, args []string) {
		run(args)
	},
}

func run(args []string) int {
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

	return 0
}
