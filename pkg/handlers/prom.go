package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.papegames.com/fengche/yayagf/pkg/prom"
)

func MountPromHandlerToGin(r gin.IRouter, options... prometheus.Collector) {
	reg := prometheus.NewRegistry()

	reg.MustRegister(prom.Routine)
	reg.MustRegister(prom.Core)
	reg.MustRegister(prom.CPU)
	reg.MustRegister(prom.Mem)
	reg.MustRegister(prom.Disk)
	reg.MustRegister(prom.Load)

	reg.MustRegister(prom.OpenConnections)
	reg.MustRegister(prom.Idle)
	reg.MustRegister(prom.InUse)
	reg.MustRegister(prom.WaitCount)
	reg.MustRegister(prom.WaitDuration)

	for _, o := range options{
		reg.MustRegister(o)
	}

	Handlers{Handler{
		path:    "/",
		handler: promhttp.InstrumentMetricHandler(
			reg, promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
		),
	}}.MountToEndpoint(r)
}
