package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MountPromHandlerToGin(r gin.IRouter, options ...prometheus.Collector) {
	reg := prometheus.NewRegistry()

	for _, o := range options {
		reg.MustRegister(o)
	}

	Handlers{Handler{
		path: "/",
		handler: promhttp.InstrumentMetricHandler(
			reg, promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
		),
	}}.MountToEndpoint(r)
}
