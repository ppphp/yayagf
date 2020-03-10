package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type Handlers []Handler

type GinRouter interface  {GET(string, ...gin.HandlerFunc)}

func (h Handlers) MountToEndpoint(r GinRouter) {
	for _, s := range h {
		r.GET(s.GetPath(), s.GetGinHandler())
		if filepath.Clean(s.GetPath()) == "/index.html" {
			r.GET("/", s.GetGinHandler())
		}
	}
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
