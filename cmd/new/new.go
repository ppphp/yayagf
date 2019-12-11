package new

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"gitlab.papegames.com/fengche/yayagf/internal/util"

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
	dir, file := filepath.Split(args[0])
	name, err := filepath.Abs(filepath.Clean(file))
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	if err := os.Mkdir(name, 0744); err != nil {
		fmt.Println(err.Error())
		return 1
	}
	if err := os.Chdir(name); err != nil {
		fmt.Println(err.Error())
		return 1
	}

	cmd := exec.Command("go", "mod", "init", filepath.Join(dir, file))
	cmd.Run()

	if err := os.Mkdir(filepath.Join(name, "app"), 0744); err != nil {
		fmt.Println(err.Error())
		return 1
	}

	if f, err := os.OpenFile("main.go", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644); err != nil {
		fmt.Println(err.Error())
		return 1
	} else {
		f.Write([]byte(fmt.Sprintf(`
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
`, filepath.Join(dir, file))))
	}
	if err := os.Mkdir(filepath.Join(name, "app", "router"), 0744); err != nil {
		fmt.Println(err.Error())
		return 1
	}
	if f, err := os.OpenFile(filepath.Join(name, "app", "router", "router.go"), os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644); err != nil {
		fmt.Println(err.Error())
		return 1
	} else {
		f.Write([]byte(fmt.Sprintf(`
package router

import (
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.Engine) {
}
`)))
	}
	util.CreateDir("app/swagger", false)
	os.Chdir(filepath.Join("app/swagger"))

	{
		cmd := exec.Command("swagger", "init", "spec")
		cmd.Run()
	}

	return 0
}

func CommandFactory() (cli.Command, error) {
	c := &Command{}
	return c, nil
}
