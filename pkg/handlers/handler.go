package handlers

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Handlers []Handler

func (h Handlers) MountToEndpoint(r gin.IRouter) {
	for _, s := range h {
		r.GET(s.GetPath(), s.GetGinHandler())
		if filepath.Clean(s.GetPath()) == "/index.html" {
			r.GET("/", s.GetGinHandler())
		}
	}
}

// pprof prom meta
func MountALotOfThingToEndpoint(r gin.IRouter, options ...prometheus.Collector) {
	MountPProfHandlerToGin(r.Group("/pprof"))
	MountMetaHandlerToGin(r.Group("/meta"))
	MountPromHandlerToGin(r.Group("/metrics"), options...)
	MountHealthHandlerToGin(r.Group("/health"))
}

type Handler struct {
	path    string
	handler http.Handler
}

func (h Handler) GetPath() string {
	return h.path
}

func (h Handler) GetHTTPHandler() http.Handler {
	return h.handler
}

func (h Handler) GetGinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		h.handler.ServeHTTP(c.Writer, c.Request)
	}
}
