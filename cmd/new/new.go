package new

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gitlab.papegames.com/fengche/yayagf/internal/swagger"

	"gitlab.papegames.com/fengche/yayagf/internal/command"
	"gitlab.papegames.com/fengche/yayagf/pkg/cli"
)

func CommandFactory() (*cli.Command, error) {
	c := &cli.Command{
		Run: func(args []string) (int, error) {
			if len(args) == 0 {
				fmt.Println("no project name")
				return 1, nil
			}
			namespace, name := filepath.Split(args[0])
			mod := filepath.Join(namespace, name)
			dir, err := filepath.Abs(filepath.Clean(name))
			if err != nil {
				fmt.Println(err.Error())
				return 1, err
			}
			log.Printf("create %v", dir)
			if err := os.MkdirAll(dir, 0644); err != nil {
				fmt.Println(err.Error())
				return 1, err
			}
			log.Printf("chdir %v", dir)
			if err := os.Chdir(dir); err != nil {
				fmt.Println(err.Error())
				return 1, err
			}

			log.Printf("init mod")
			command.DoCommand("go", []string{"mod", "init", mod}, nil, nil)

			log.Printf("create %v", filepath.Join(dir, "main.go"))
			ioutil.WriteFile(filepath.Join(dir, "main.go"), []byte(fmt.Sprintf(`
package main

import (
	"github.com/gin-contrib/cors"
	// "gitlab.papegames.com/fengche/yayagf/pkg/model"
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

	// Uncomment the following code to simplify db
	//drv, err := model.Open("mysql", config.GetConfig().DB)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//crud.C = crud.NewClient(crud.Driver(drv))
	//if err := crud.C.Schema.Create(context.Background()); err != nil {
	//	log.Fatal(err)
	//}

	if err := r.Run(fmt.Sprintf(":%v", config.GetConfig().Port)); err != nil {
		log.Fatal(err)
	}
}
`, mod, mod, "%v")), 0644)
			log.Printf("create %v", filepath.Join(dir, "app", "router", "router.go"))
			if err := ioutil.WriteFile(filepath.Join(dir, "app", "router", "router.go"), []byte(fmt.Sprintf(`
package router

import (
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.Engine) {
}
`)), 0644); err != nil {
				log.Println(err.Error())
				return 1, err
			}
			if err := ioutil.WriteFile(filepath.Join(dir, "app", "config", "config.go"), []byte(`
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

`), 0644); err != nil {
				log.Println(err.Error())
				return 1, err
			}
			if err := ioutil.WriteFile(filepath.Join(dir, "conf.toml"), []byte(`
db=""
port=8080
`), 0644); err != nil {
				log.Println(err.Error())
				return 1, err
			}

			out, errs := &bytes.Buffer{}, &bytes.Buffer{}
			log.Printf("init swagger")
			if err := swagger.GenerateSwagger(); err != nil {
				log.Fatalf("swag failed %v", errs.String())
				return 1, err
			}

			log.Printf("init git")
			if err := command.DoCommand("git", []string{"init"}, out, errs); err != nil {
				log.Fatalf("git failed %v", errs.String())
				return 1, err
			}
			if err := ioutil.WriteFile(filepath.Join(dir, ".gitignore"), []byte(fmt.Sprintf(`
%v
%v.tar
`, name, name)), 0644); err != nil {
				log.Fatalf("gitignore failed %v", errs.String())
				return 1, err
			}

			log.Printf("init docker")
			if err := ioutil.WriteFile(filepath.Join(dir, "Dockerfile"), []byte(fmt.Sprintf(`
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

`)), 0644); err != nil {
				log.Fatalf("docker failed %v", errs.String())
				return 1, err
			}

			return 0, nil
		},
	}
	return c, nil
}
