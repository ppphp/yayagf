package new

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gitlab.papegames.com/fengche/yayagf/internal/blueprint"
	"gitlab.papegames.com/fengche/yayagf/internal/command"
	"gitlab.papegames.com/fengche/yayagf/internal/log"
	"gitlab.papegames.com/fengche/yayagf/internal/swagger"
	"gitlab.papegames.com/fengche/yayagf/pkg/cli"
)

func CommandFactory() (*cli.Command, error) {
	return &cli.Command{Run: runNew}, nil
}

func runNew(args []string, flags map[string]string) (int, error) {
	if len(args) == 0 {
		log.Errorf("no project name")
		return 1, nil
	}
	namespace, name := filepath.Split(args[0])
	mod := filepath.Join(namespace, name)
	dir, err := filepath.Abs(filepath.Clean(name))
	if err != nil {
		log.Errorf("abs(%v) failed %v", name, err)
		return 1, err
	}
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Errorf("create (%v) failed %v", dir, err)
		return 1, err
	}
	if err := os.Chdir(dir); err != nil {
		log.Errorf("chdir (%v) failed %v", dir, err)
		return 1, err
	}

	log.Printf("init project")
	if err, o, e := command.DoCommand("go", "mod", "init", mod); err != nil {
		log.Errorf("go mod failed %v %v\n", o, e)
		return 1, err
	}

	log.Printf("create %v", filepath.Join(dir, "main.go"))
	if err := blueprint.WriteFileWithTmpl(filepath.Join(dir, "main.go"), MainGo, struct{ Mod string }{mod}); err != nil {
		log.Println(err.Error())
		return 1, err
	}
	log.Printf("create %v", filepath.Join(dir, "app", "router", "router.go"))
	if err := blueprint.WriteFileWithTmpl(filepath.Join(dir, "app", "router", "router.go"), RouterGo, nil); err != nil {
		log.Println(err.Error())
		return 1, err
	}

	if err := blueprint.WriteFileWithTmpl(filepath.Join(dir, "app", "config", "config.go"), ConfigGo, nil); err != nil {
		log.Println(err.Error())
		return 1, err
	}
	if err := blueprint.WriteFileWithTmpl(filepath.Join(dir, "conf.toml"), fmt.Sprintf(`
db=""
port=8080
[log]
filename="./log/%v.log"
level=5
`, name), nil); err != nil {
		log.Println(err.Error())
		return 1, err
	}

	log.Printf("init swagger")
	if err := swagger.GenerateSwagger(); err != nil {
		log.Errorf("swag failed %v", err)
		return 1, err
	}

	log.Printf("init git")
	if err, _, errs := command.DoCommand("git", "init"); err != nil {
		log.Errorf("git failed %v", errs)
		return 1, err
	}
	if err := blueprint.WriteFileWithTmpl(filepath.Join(dir, ".gitignore"), `
{{.Name}}
{{.Name}}.tar
`, struct{ Name string }{name}); err != nil {
		log.Errorf("gitignore failed %v", err)
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
		log.Errorf("docker failed %v", err)
		return 1, err
	}

	if err := ioutil.WriteFile(filepath.Join(dir, "README.md"), []byte("\n"), 0644); err != nil {
		log.Errorf("readme failed %v", err)
		return 1, err
	}

	return 0, nil
}

const (
	MainGo = `
package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	// {{.Mod}}/app/crud"
	"gitlab.papegames.com/fengche/yayagf/pkg/handlers"
	"gitlab.papegames.com/fengche/yayagf/pkg/log"
	// "gitlab.papegames.com/fengche/yayagf/pkg/model"
	"gitlab.papegames.com/fengche/yayagf/pkg/prom"
	"github.com/gin-gonic/gin"
	"{{.Mod}}/app/router"
	"{{.Mod}}/app/config"
)
// @title "{{.Mod}} API
// @version master
// @description This is a {{.Mod}} server

// @contact.name 风车
// @contact.url https://{{.Mod}}
// @contact.email liukaiwen@papegames.net

// @host localhost:8080
// @BasePath /api/v1

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Errorf("load config failed (%v)", err)
		return
	}

	log.Tweak(config.GetConfig().Log)
	gin.DefaultWriter = log.GetLogger().Out
	r := maotai.Default("giftsvr")
	router.RegisterRouter(r)

	drv, err := model.Open("mysql", config.GetConfig().DB)
	if err != nil {
		log.Errorf("create sql driver failed (%v)", err)
		return
	}
	//crud.C = crud.NewClient(crud.Driver(drv))
	//if err := crud.C.Schema.Create(context.Background()); err != nil {
	//	log.Fatal(err)
	//}
	handlers.MountALotOfThingToEndpoint(r.Group("admin"),
		handlers.WithMetric(r.TTLHist, r.URLConn,
			prom.SysCPU(), prom.SysMem(), prom.SysDisk(), prom.SysLoad(), prom.GoRoutine(), prom.GoMem(),
			prom.DbConnection(config.GetConfig().DB, drv.DB()), prom.DBWaitCount(config.GetConfig().DB, drv.DB()),
			prom.DBWaitDuration(config.GetConfig().DB, drv.DB()), prom.DbClose(config.GetConfig().DB, drv.DB()),
			prom.BuildInfo()),
		handlers.WithSwagger(doc.Swagger),
	)

	if err := r.Run(fmt.Sprintf(":%v", config.GetConfig().Port)); err != nil {
		log.Fatal(err)
	}
}
`

	RouterGo = `package router

import (
	"github.com/gin-contrib/cors"
	"gitlab.papegames.com/fengche/yayagf/pkg/log"
	"gitlab.papegames.com/fengche/yayagf/pkg/maotai"
	"gitlab.papegames.com/fengche/yayagf/pkg/middleware"
)

func RegisterRouter(r *maotai.MaoTai) {
	r.Use(cors.Default())
	api := r.Group("/")
	api.Use(middleware.Ginrus(log.GetLogger()))

}
`

	ConfigGo = `package config

import (
	"sync"

	"gitlab.papegames.com/fengche/yayagf/pkg/config"
	"gitlab.papegames.com/fengche/yayagf/pkg/log"
)

var lock sync.RWMutex

type Config struct {
	DB   string
	Port int
	Log  log.Config
}

var conf = new(Config)

// only support ini like config
func LoadConfig() error {
	lock.Lock()
	defer lock.Unlock()
	if err := config.LoadConfig("conf.toml", conf); err != nil {
		return err
	}

	log.Infof("%v", conf)

	return nil
}

func GetConfig() Config {
	lock.RLock()
	defer lock.RUnlock()
	return *conf
}
`
)
