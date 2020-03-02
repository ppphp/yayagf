package generate

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"gitlab.papegames.com/fengche/yayagf/internal/file"
)

var Command = &cobra.Command{
	Use: "generate",
	Run: func(cmd *cobra.Command, args []string) {
		run(args)
	},
}

func run(args []string) int {
	if len(args) == 0 {
		log.Println("need generate something")
		return 1
	}

	switch args[0] {
	case "table":
		pwd, err := os.Getwd()
		if err != nil {
			log.Panic(err)
		}
		root, err := file.FindAppRoot(pwd)
		if err != nil {
			log.Panic(err)
		}
		f, err := file.CreateFile(filepath.Join(root, "migrates"), false)
		if err != nil {
			log.Panic(err)
		}
		f.WriteString("")
	default:
		log.Println("need generate something")
		return 1
	}

	return 0
}
