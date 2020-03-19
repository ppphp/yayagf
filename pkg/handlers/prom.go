package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)
var PromHandler = promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			EnableOpenMetrics: true,
		},
	)

func MountPromHandlerToGin(r gin.IRouter) {
	Handlers{Handler{
		path: "/",
		handler: PromHandler,
	}}.MountToEndpoint(r)
}