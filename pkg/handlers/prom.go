package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.papegames.com/fengche/yayagf/pkg/prom"
)

func MountPromHandlerToGin(r gin.IRouter) {
	reg := prometheus.NewRegistry()

	reg.MustRegister(prom.Routine)
	reg.MustRegister(prom.Core)
	reg.MustRegister(prom.CPU)
	reg.MustRegister(prom.Mem)
	reg.MustRegister(prom.Disk)
	reg.MustRegister(prom.Load)

	Handlers{Handler{
		path:    "/",
		handler: promhttp.InstrumentMetricHandler(
			reg, promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
		),
	}}.MountToEndpoint(r)
}
