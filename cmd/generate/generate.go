package generate

import (
	"github.com/spf13/cobra"
	"gitlab.papegames.com/fengche/yayagf/cmd/model"
)

var Command = &cobra.Command{
	Use:   "generate",
	Short: "generate",
}

var AliasCommand = &cobra.Command{
	Use:   "g",
	Short: "generate",
}

func init() {
	Command.AddCommand(
		model.Command,
	)
	AliasCommand.AddCommand(
		model.Command,
	)
}
