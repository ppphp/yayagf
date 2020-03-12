package model

import (
	"bytes"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gitlab.papegames.com/fengche/yayagf/internal/command"
	"gitlab.papegames.com/fengche/yayagf/internal/file"
)

var Command = &cobra.Command{
	Use: "model",
	Run: func(cmd *cobra.Command, args []string) {
		run(args)
	},
}

func run(args []string) int {
	root, err := file.GetAppRoot()
	if err != nil {
		log.Printf("get project name failed: %v", err.Error())
		return 1
	}
	if err := os.Chdir(filepath.Join(root, "app")); err != nil {
		log.Printf("chdir failed: %v", err.Error())
		return 1
	}
	out, errs := &bytes.Buffer{}, &bytes.Buffer{}
	if err := command.DoCommand("entc", []string{"init"}, out, errs); err != nil {
		log.Fatalf("ent init failed: %v", errs.String())
		return 1
	}

	return 0
}
