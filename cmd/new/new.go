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

	log.Printf("create %v", filepath.Join(dir, "main.go"))
	file.CreateFileWithContent(filepath.Join(dir, "main.go"), fmt.Sprintf(`
package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"%v/app/router"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	r := gin.Default()

	r.Use(cors.Default())

	router.AddRoute(r)

	r.Run()
}
`, mod))
	log.Printf("create %v", filepath.Join(dir, "app", "router", "router.go"))
	if err := file.CreateFileWithContent(filepath.Join(dir, "app", "router", "router.go"), fmt.Sprintf(`
package router

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/gin-gonic/gin"
	_ "%v/app/docs"
)

func AddRoute(r *gin.Engine) {
	url := ginSwagger.URL("http://localhost:3000/v1/docs/doc.json") // The url pointing to API definition
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
`), mod); err != nil {
		log.Println(err.Error())
		return 1
	}

	log.Printf("init swagger")
	command.DoCommand("swag", []string{"init", "spec", "-o", "app/docs"}, nil, nil)

	log.Printf("init git")
	command.DoCommand("git", []string{"init"}, nil, nil)

	return 0
}

func CommandFactory() (cli.Command, error) {
	c := &Command{}
	return c, nil
}
