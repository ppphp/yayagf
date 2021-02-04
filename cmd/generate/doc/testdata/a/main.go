
package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	// a/app/crud"
	"gitlab.papegames.com/fengche/yayagf/pkg/handlers"
	"gitlab.papegames.com/fengche/yayagf/pkg/log"
	// "gitlab.papegames.com/fengche/yayagf/pkg/model"
	"gitlab.papegames.com/fengche/yayagf/pkg/prom"
	"github.com/gin-gonic/gin"
	"a/app/router"
	"a/app/config"
)
// @title "a API
// @version master
// @description This is a a server

// @contact.name 风车
// @contact.url https://a
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
