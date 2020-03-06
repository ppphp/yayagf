package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
