package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MountPromHandlerToGin(r gin.IRouter) {
	Handlers{Handler{
		path: "/",
		handler: promhttp.Handler(),
	}}.MountToEndpoint(r)
}