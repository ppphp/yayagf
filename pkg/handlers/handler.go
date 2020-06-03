package handlers

import (
	"net/http"
	"path/filepath"

	"github.com/prometheus/client_golang/prometheus"

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

type MountOption struct {
	metric []prometheus.Collector
}

func WithMetric(collectors ...prometheus.Collector) *MountOption {
	return &MountOption{metric: collectors}
}

// pprof prom meta
func MountALotOfThingToEndpoint(r gin.IRouter, options ...*MountOption) {
	MountPProfHandlerToGin(r.Group("/pprof"))
	MountMetaHandlerToGin(r.Group("/meta"))
	cs := []prometheus.Collector{}
	for _, o := range options {
		cs = append(cs, o.metric...)
	}
	MountPromHandlerToGin(r.Group("/metrics"), cs...)
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
