package interactive

import (
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use: "interactive",
	Run: func(cmd *cobra.Command, args []string) {
		run(args)
	},
}

func run(args []string) int {
	panic("WIP")
	return 0
}
