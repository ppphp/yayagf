# yayagf

yet another yet another go web framework
[![pipeline status](https://gitlab.papegames.com/fengche/yayagf/badges/master/pipeline.svg)](https://gitlab.papegames.com/fengche/yayagf/commits/master)
[![coverage report](https://gitlab.papegames.com/fengche/yayagf/badges/master/coverage.svg)](https://gitlab.papegames.com/fengche/yayagf/commits/master)

## install

`go install ./cmd/yayagf`

## what to use

### performance

web framework should not treat the performance very important.

go is fast, so performance is not a really big thing.

util now, it uses gin as framework.

## how to use

### yayagf new

type `yayagf new gitlab.papegames.com/fengche/abc` will generate a folder named abc

a project structure is like this

project_root
-app
--controller
--crud
--schema
--serializer // TODO
--docs // TODO
-config_file1 // TODO
-config_file2 // TODO
-project_root.yml // TODO
-Dockerfile // TODO

### yayagf server // FIXME

two steps

1. go into any go project 

2. `yayagf server` will monitor the code, recompile the code when compilable and run.

### yayagf package

package to docker // TODO some others

### yayagf generate

generate a component scaffold.

```bash
yayagf g commands
yayagf generate commands
```

#### schema (ent)

will generate schema in `app/schema`

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

will generate crud files in `app/crud` according to `app/schema`

```bash
yayagf g doc
```


#### router // TODO

#### serializer // TODO

## packages

### ent driver

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