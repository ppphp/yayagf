package _package

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/ppphp/yayagf/internal/command"
	"github.com/ppphp/yayagf/internal/file"
	"github.com/ppphp/yayagf/pkg/cli"
)

func CommandFactory() (*cli.Command, error) {
	c := &cli.Command{
		Run: func(args []string, flags map[string]string) (int, error) {
			root, err := file.GetAppRoot()
			name := filepath.Base(root)
			if err != nil {
				log.Printf("get project name failed: %v", err.Error())
				return 1, err
			}

			priB, err := ioutil.ReadFile(os.Getenv("HOME") + "/.ssh/id_rsa")
			if err != nil {
				log.Println(err)
			}
			priS := string(priB)
			os.Setenv("pri", priS)

			pubB, err := ioutil.ReadFile(os.Getenv("HOME") + "/.ssh/id_rsa.pub")
			if err != nil {
				log.Println(err)
			}
			pubS := string(pubB)
			os.Setenv("pub", pubS)

			if err, _, errs := command.DoCommand("docker", "build", "-t", fmt.Sprintf("docker.papegames.com/%v", name), ".", "--build-arg", "pri", "--build-arg", "pub"); err != nil {
				log.Fatalf("docker build failed: %v", errs)
				return 1, err
			}
			if err, _, errs := command.DoCommand("docker", "save", fmt.Sprintf("docker.papegames.com/%v", name), "-o", fmt.Sprintf("%v.tar", name)); err != nil {
				log.Fatalf("docker save error: %v", errs)
				return 1, err
			}

			return 0, nil
		},
	}
	return c, nil
}
