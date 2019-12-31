package new

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gitlab.papegames.com/fengche/yayagf/internal/command"
	"gitlab.papegames.com/fengche/yayagf/internal/file"

	"github.com/mitchellh/cli"
)

type Command struct {
}

func (c *Command) Help() string {
	return ""
}

func (c *Command) Synopsis() string {
	return "init a yayagf project"
}

func (c *Command) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println("no project name")
		return 1
	}
	namespace, name := filepath.Split(args[0])
	mod := filepath.Join(namespace, name)
	dir, err := filepath.Abs(filepath.Clean(name))
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	log.Printf("create %v", dir)
	if err := file.CreateDir(dir, false); err != nil {
		fmt.Println(err.Error())
		return 1
	}
	log.Printf("chdir %v", dir)
	if err := os.Chdir(dir); err != nil {
		fmt.Println(err.Error())
		return 1
	}

	log.Printf("init mod")
	command.DoCommand("go", []string{"mod", "init", mod}, nil, nil)

	log.Printf("create %v", filepath.Join(dir, "app", "main.go"))
	file.CreateFileWithContent(filepath.Join(dir, "app", "main.go"), fmt.Sprintf(`
package main

import (
	"github.com/gin-gonic/gin"
	"%v/app/router"
)

func main() {
	r := gin.Default()

	router.AddRoute(r)

	r.Run()
}
`, mod))
	log.Printf("create %v", filepath.Join(dir, "app", "router", "router.go"))
	if err := file.CreateFileWithContent(filepath.Join(dir, "app", "router", "router.go"), fmt.Sprintf(`
package router

import (
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.Engine) {
}
`)); err != nil {
		log.Println(err.Error())
		return 1
	}

	log.Printf("init swagger")
	command.DoCommand("swagger", []string{"init", "spec"}, nil, nil)

	log.Printf("init git")
	command.DoCommand("git", []string{"init"}, nil, nil)

	return 0
}

func CommandFactory() (cli.Command, error) {
	c := &Command{}
	return c, nil
}
