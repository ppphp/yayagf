package router

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
