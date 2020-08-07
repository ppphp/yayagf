
[![pipeline status](https://gitlab.papegames.com/fengche/yayagf/badges/master/pipeline.svg)](https://gitlab.papegames.com/fengche/yayagf/commits/master)
[![coverage report](https://gitlab.papegames.com/fengche/yayagf/badges/master/coverage.svg)](https://gitlab.papegames.com/fengche/yayagf/commits/master)
[![SQALE评级](http://192.168.0.97:9000/api/project_badges/measure?project=yayagf&metric=sqale_rating)](http://192.168.0.97:9000/dashboard?id=yayagf)
# yayagf

yet another yet another go web framework, my practice

it is a 缝合怪 of some tool
- http: gin-gonic/gin
- monitor: prometheus/prometheus
- database: facebookincubator/ent
- cli: urfave/cli (WIP)

###### sonar
[![Bugs](http://192.168.0.97:9000/api/project_badges/measure?project=yayagf&metric=bugs)](http://192.168.0.97:9000/dashboard?id=yayagf)
[![异味](http://192.168.0.97:9000/api/project_badges/measure?project=yayagf&metric=code_smells)](http://192.168.0.97:9000/dashboard?id=yayagf)
[![覆盖率](http://192.168.0.97:9000/api/project_badges/measure?project=yayagf&metric=coverage)](http://192.168.0.97:9000/dashboard?id=yayagf)
[![重复行(%)](http://192.168.0.97:9000/api/project_badges/measure?project=yayagf&metric=duplicated_lines_density)](http://192.168.0.97:9000/dashboard?id=yayagf)
[![代码行数](http://192.168.0.97:9000/api/project_badges/measure?project=yayagf&metric=ncloc)](http://192.168.0.97:9000/dashboard?id=yayagf)
[![警报](http://192.168.0.97:9000/api/project_badges/measure?project=yayagf&metric=alert_status)](http://192.168.0.97:9000/dashboard?id=yayagf)
[![可靠性比率](http://192.168.0.97:9000/api/project_badges/measure?project=yayagf&metric=reliability_rating)](http://192.168.0.97:9000/dashboard?id=yayagf)
[![安全比率](http://192.168.0.97:9000/api/project_badges/measure?project=yayagf&metric=security_rating)](http://192.168.0.97:9000/dashboard?id=yayagf)
[![技术债务](http://192.168.0.97:9000/api/project_badges/measure?project=yayagf&metric=sqale_index)](http://192.168.0.97:9000/dashboard?id=yayagf)
[![漏洞](http://192.168.0.97:9000/api/project_badges/measure?project=yayagf&metric=vulnerabilities)](http://192.168.0.97:9000/dashboard?id=yayagf)


## install

`go install ./cmd/yayagf`

## start a project

type `yayagf new gitlab.papegames.com/fengche/abc` will generate a folder named abc

a project structure is like this

project_root
-app
--controller
--crud
--schema
--serializer // TODO
--doc
-config_file1 // TODO
-config_file2 // TODO
-project_root.yml // TODO
-Dockerfile

## run a project

`yayagf server` will monitor the code, recompile the code when compilable and run.

## generate a component scaffold

```bash
yayagf g commands
yayagf generate commands
```

#### schema (ent)

will generate schema(table) in `app/schema`

```bash
yayagf g schema a
```

#### crud (ent)

will generate crud files in `app/crud` according to `app/schema`

```bash
yayagf g curd
yayagf g crud
```

#### doc (swagger)

will generate doc files in `app/doc` according to `app/controllers`

```bash
yayagf g doc
```


#### router // TODO

#### serializer // TODO

## write codes

### ent (db)
- driver
```go
package main
import (
    "gitlab.papegames.com/fengche/yayagf/pkg/model"
    "log"
    "yourproject/ent"
)

func main() {
	drv, err := model.Open("mysql", DBURL)
	if err != nil {
		log.Fatal(drv)
	}
	client := ent.NewClient(ent.Driver(drv))
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
	}
}
```
- sharding name
- sharding databases
- monitoring


### config loader

```go
package main
import (
    "gitlab.papegames.com/fengche/yayagf/pkg/config"
    "log"
)

func main() {
    var conf struct {
        A int
        B string
    }
    config.LoadConfig(&conf)
    log.Println(conf)
}
```

### util handlers
```go
package main
import (
    "gitlab.papegames.com/fengche/yayagf/pkg/handlers"
    "log"
)

func main() {
    // a static router
	if ss, err := handlers.ServeStaticDirectory(config.GetConfig().Static); err != nil {
		log.Println(err)
	} else {
		for _, s := range ss {
			r.GET(s.GetPath(), s.GetGinHandler())
			if filepath.Clean(s.GetPath()) == "/index.html" {
				r.GET("/", s.GetGinHandler())
			}
		}
	}

	// a pprof generator
	for _, s := range handlers.PProfHandlers {
		r.Group("pprof").GET(s.GetPath(), s.GetGinHandler())
	}
}
```

### monitoring

only monitor dynamic metrics

#### system
- cpu
- memory
- load average
- disk

#### runtime
- process cpu
- process mem
- process fd
- thread
- goroutine
- gc

#### handler
- outer services (rest)
- inner services (rpc)

#### storage
- db
- redis
- thirdparty services (rpc)

## testing

WIP

## performance

i don't care.

## packages

package to docker // TODO some others
