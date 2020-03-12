package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.papegames.com/fengche/yayagf/cmd/doc"
	"gitlab.papegames.com/fengche/yayagf/cmd/generate"
	"gitlab.papegames.com/fengche/yayagf/cmd/interactive"
	"gitlab.papegames.com/fengche/yayagf/cmd/model"
	"gitlab.papegames.com/fengche/yayagf/cmd/new"
	_package "gitlab.papegames.com/fengche/yayagf/cmd/package"
	"gitlab.papegames.com/fengche/yayagf/cmd/server"
)

var RootCmd = &cobra.Command{
	Short: "yet another yet another go framework cli interface",
	Long:  `yet another yet another go framework cli interface just like rails`,
}

func init() {
	RootCmd.AddCommand(
		new.Command,
		server.Command,
		interactive.Command,
		_package.Command,
		generate.Command, generate.AliasCommand,
		model.Command,
		doc.Command,
	)
}

// Execute executes the root command.
func Execute() error {
	return RootCmd.Execute()
}
