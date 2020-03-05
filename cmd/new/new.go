package new

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"gitlab.papegames.com/fengche/yayagf/internal/command"
	"gitlab.papegames.com/fengche/yayagf/internal/file"
)

var Command = &cobra.Command{
	Use: "new",
	Run: func(cmd *cobra.Command, args []string) {
		run(args)
	},
}

func run(args []string) int {
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
	"log"
	"%v/app/router"
	"%v/app/config"
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

	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	r.Run()
}
`, mod, mod))
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
`, mod)); err != nil {
		log.Println(err.Error())
		return 1
	}
	if err := file.CreateFileWithContent(filepath.Join(dir, "app", "config", "config.go"), `
package config

import (
	"gitlab.papegames.com/fengche/yayagf/pkg/config"
	"log"
	"sync"
)

var lock sync.RWMutex

type Config struct {
	DB   string
	Port int
}

var conf = new(Config)

// only support ini like config
func LoadConfig() error {
	lock.Lock()
	defer lock.Unlock()
	if err := config.LoadTomlFile("conf.toml", conf); err != nil {
		log.Fatal(err)
	}

	config.LoadEnv(conf)

	log.Println(conf)
	return nil
}

func GetConfig() Config {
	lock.RLock()
	defer lock.RUnlock()
	return *conf
}

`); err != nil {
		log.Println(err.Error())
		return 1
	}
	if err := file.CreateFileWithContent(filepath.Join(dir, "conf.toml"), `
db=""
port=8080
`); err != nil {
		log.Println(err.Error())
		return 1
	}

	out, errs := &bytes.Buffer{}, &bytes.Buffer{}
	log.Printf("init swagger")
	if err := command.DoCommand("swag", []string{"init", "spec", "-o", "app/docs"}, out, errs); err != nil {
		log.Fatalf("swag failed %v", errs.String())
		return 1
	}

	log.Printf("init git")
	if err := command.DoCommand("git", []string{"init"}, out, errs); err != nil {
		log.Fatalf("git failed %v", errs.String())
		return 1
	}
	if err := file.CreateFileWithContent(filepath.Join(dir, ".gitignore"), fmt.Sprintf(`
%v
%v.tar
`, name, name)); err != nil {
		log.Fatalf("gitignore failed %v", errs.String())
		return 1
	}

	log.Printf("init docker")
	if err := file.CreateFileWithContent(filepath.Join(dir, "Dockerfile"), fmt.Sprintf(`
FROM golang as back

ENV GOPROXY=https://goproxy.io
ENV GOSUMDB=off
ENV GOPRIVATE=gitlab.papegames.com/*
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on

WORKDIR /main

ADD go.mod ./
ADD go.sum ./
RUN go mod download

ADD app ./app/
ADD main.go ./
RUN go build -o /main/main

FROM scratch
WORKDIR /main
COPY --from=back /main/main .

CMD ["/main/main"]

`, )); err != nil {
		log.Fatalf("docker failed %v", errs.String())
		return 1
	}

	log.Printf("init ent")
	if err := os.Chdir(filepath.Join(dir, "app")); err != nil {
		log.Printf("chdir failed: %v", err.Error())
		return 1
	}
	if err := command.DoCommand("entc", []string{"init",}, out, errs); err != nil {
		log.Fatalf("ent init failed: %v", errs.String())
		return 1
	}
	return 0
}
