package model

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.papegames.com/fengche/yayagf/internal/command"
	"gitlab.papegames.com/fengche/yayagf/internal/file"
	"log"
	"path/filepath"
)

var Command = &cobra.Command{
	Use: "model",
	Run: func(cmd *cobra.Command, args []string) {
		run(args)
	},
}

func run(args []string) int {
	root, err := file.GetAppRoot()
	name := filepath.Base(root)
	if err != nil {
		log.Printf("get project name failed: %v", err.Error())
		return 1
	}
	out, errs := &bytes.Buffer{}, &bytes.Buffer{}
	if err := command.DoCommand("docker", []string{"build", "-t", fmt.Sprintf("docker.papegames.com/%v", name), "."}, out, errs); err != nil {
		log.Fatalf("docker build failed: %v", errs.String())
		return 1
	}
	if err := command.DoCommand("docker", []string{"save", fmt.Sprintf("docker.papegames.com/%v", name), "-o", fmt.Sprintf("%v.tar", name)}, out, errs); err != nil {
		log.Fatalf("docker save error: %v", errs.String())
		return 1
	}

	return 0
}
