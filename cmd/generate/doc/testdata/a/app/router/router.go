package router

import (
	"github.com/gin-contrib/cors"
	"github.com/ppphp/yayagf/pkg/log"
	"github.com/ppphp/yayagf/pkg/maotai"
	"github.com/ppphp/yayagf/pkg/middleware"
)

func RegisterRouter(r *maotai.MaoTai) {
	r.Use(cors.Default())
	api := r.Group("/")
	api.Use(middleware.Ginrus(log.GetLogger()))

}
